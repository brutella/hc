package hap

import (
	"bytes"
	"io/ioutil"
	"net"
	"sync"

	"github.com/brutella/hc/db"
	"github.com/brutella/hc/event"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/container"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/server"
	"github.com/brutella/hc/util"
	"github.com/brutella/log"
	"github.com/gosexy/to"
)

// Config provides basic configuration for an IP transport
type Config struct {
	// Path to the storage
	// When empty, the tranport stores the data inside a folder named exactly like the accessory
	StoragePath string

	// Port at which transport is reachable e.g. 12345
	// When empty, the transport uses a random port
	Port string

	// Pin with has to be entered on iOS client to pair with the accessory
	// When empty, the pin 00102003 is used
	Pin string
}

type ipTransport struct {
	config  Config
	context netio.HAPContext
	server  server.Server
	mutex   *sync.Mutex
	mdns    *MDNSService

	storage  util.Storage
	database db.Database

	name      string
	device    netio.SecuredDevice
	container *container.Container

	// Used to communicate between different parts of the program (e.g. successful pairing with HomeKit)
	emitter event.Emitter
}

// NewIPTransport creates a transport to provide accessories over IP.
//
// The IP transports stores the crypto keys inside a database, which
// is by default inside a folder at the current working directory.
// The folder is named exactly as the accessory name.
//
// The transports can contain more than one accessory. If this is the
// case, the first accessory acts as the HomeKit bridge.
//
// *Important:* Changing the name of the accessory, or letting multiple
// transports store the data inside the same database lead to
// unexpected behavior – don't do that.
//
// The transport is secured with an 8-digit pin, which must be entered
// by an iOS client to successfully pair with the accessory. If the 
// provided transport config does not specify any pin, 00102003 is used.
func NewIPTransport(config Config, a *accessory.Accessory, as ...*accessory.Accessory) (Transport, error) {
	// Find transport name which is visible in mDNS
	name := a.Name()
	if len(name) == 0 {
		log.Fatal("Invalid empty name for first accessory")
	}

	default_config := Config{
		StoragePath: name,
		Pin:         "00102003",
		Port:        "",
	}

	if dir := config.StoragePath; len(dir) > 0 {
		default_config.StoragePath = dir
	}

	if pin := config.Pin; len(pin) > 0 {
		default_config.Pin = pin
	}

	if port := config.Port; len(port) > 0 {
		default_config.Port = ":" + port
	}

	storage, err := util.NewFileStorage(default_config.StoragePath)
	if err != nil {
		return nil, err
	}

	// Find transport uuid which appears as "id" txt record in mDNS and
	// must be unique and stay the same over time
	uuid := transportUUIDInStorage(storage)
	database := db.NewDatabaseWithStorage(storage)

	hap_pin, err := NewPin(default_config.Pin)
	if err != nil {
		return nil, err
	}

	device, err := netio.NewSecuredDevice(uuid, hap_pin, database)

	t := &ipTransport{
		database:  database,
		name:      name,
		device:    device,
		config:    default_config,
		container: container.NewContainer(),
		mutex:     &sync.Mutex{},
		context:   netio.NewContextForSecuredDevice(device),
		emitter:   event.NewEmitter(),
	}

	t.addAccessory(a)
	for _, a := range as {
		t.addAccessory(a)
	}

	t.emitter.AddListener(t)

	return t, err
}

func (t *ipTransport) Start() {
	config := server.Config{
		Port:      t.config.Port,
		Context:   t.context,
		Database:  t.database,
		Container: t.container,
		Device:    t.device,
		Mutex:     t.mutex,
		Emitter:   t.emitter,
	}

	s := server.NewServer(config)
	t.server = s
	port := to.Int64(s.Port())

	mdns := NewMDNSService(t.name, t.device.Name(), int(port))
	t.mdns = mdns

	// Paired accessories must not be reachable for other clients since iOS 9
	if t.isPaired() {
		mdns.SetReachable(false)
	}

	mdns.Publish()

	// Listen until server.Stop() is called
	s.ListenAndServe()
}

// Stop stops the ip transport by unpublishing the mDNS service.
func (t *ipTransport) Stop() {
	if t.mdns != nil {
		t.mdns.Stop()
	}

	if t.server != nil {
		t.server.Stop()
	}
}

// isPaired returns true when the transport is already paired
func (t *ipTransport) isPaired() bool {

	// If more than one entity is stored in the database, we are paired with a device.
	// The transport itself is a device and is stored in the database, therefore
	// we have to check for more than one entity.
	if es, err := t.database.Entities(); err == nil && len(es) > 1 {
		return true
	}

	return false
}

func (t *ipTransport) updateMDNSReachability() {
	if mdns := t.mdns; mdns != nil {
		mdns.SetReachable(t.isPaired() == false)
		mdns.Update()
	}
}

func (t *ipTransport) addAccessory(a *accessory.Accessory) {
	t.container.AddAccessory(a)

	for _, s := range a.Services {
		for _, c := range s.Characteristics {
			// When a characteristic value changes and events are enabled for this characteristic
			// all listeners are notified. Since we don't track which client is interested in
			// which characteristic change event, we send them to all active connections.
			onConnChange := func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
				if c.Events == true {
					t.notifyListener(a, c, conn)
				}
			}
			c.OnConnChange(onConnChange)

			onChange := func(c *characteristic.Characteristic, new, old interface{}) {
				if c.Events == true {
					t.notifyListener(a, c, nil)
				}
			}
			c.OnChange(onChange)
		}
	}
}

func (t *ipTransport) notifyListener(a *accessory.Accessory, c *characteristic.Characteristic, except net.Conn) {
	conns := t.context.ActiveConnections()
	for _, conn := range conns {
		if conn == except {
			continue
		}
		resp, err := netio.New(a, c)
		if err != nil {
			log.Fatal(err)
		}

		// Write response into buffer to replace HTTP protocol
		// specifier with EVENT as required by HAP
		var buffer = new(bytes.Buffer)
		resp.Write(buffer)
		bytes, err := ioutil.ReadAll(buffer)
		bytes = netio.FixProtocolSpecifier(bytes)
		log.Printf("[VERB] %s <- %s", conn.RemoteAddr(), string(bytes))
		conn.Write(bytes)
	}
}

// transportUUIDInStorage returns the uuid stored in storage or
// creates a new random uuid and stores it.
func transportUUIDInStorage(storage util.Storage) string {
	uuid, err := storage.Get("uuid")
	if len(uuid) == 0 || err != nil {
		str := util.RandomHexString()
		uuid = []byte(netio.MAC48Address(str))
		err := storage.Set("uuid", uuid)
		if err != nil {
			log.Fatal(err)
		}
	}
	return string(uuid)
}

// Handles event which are sent when pairing with a device is added or removed
func (t *ipTransport) Handle(ev interface{}) {
	switch ev.(type) {
	case event.DevicePaired:
		log.Printf("[INFO] Event: paired with device")
		t.updateMDNSReachability()
	case event.DeviceUnpaired:
		log.Printf("[INFO] Event: unpaired with device")
		t.updateMDNSReachability()
	default:
		break
	}
}
