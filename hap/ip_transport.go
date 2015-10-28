package hap

import (
	"bytes"
	"io/ioutil"
	"net"
        "path"
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

type ipTransport struct {
	context netio.HAPContext
	server  server.Server
	mutex   *sync.Mutex
	mdns    *MDNSService

	storage  util.Storage
	database db.Database

	name      string
	device    netio.SecuredDevice
	container *container.Container

	emitter event.Emitter

        config *HAPConfig
}
type HAPConfig struct {
        PIN string
        StoragePath string
}

// NewIPTransport is a Compatibility wrapper for any code that depends on the simple implementation
func NewIPTransport(pin string, a *accessory.Accessory, as ...*accessory.Accessory) (Transport, error) {
        config := &HAPConfig{PIN: pin}
        return NewHAPTransport(config, a, as...)
}

// NewHAPTransport creates a transport to provide accessories over IP.
// The pairing is secured using a 8-numbers pin.
// If more than one accessory is provided, the first becomes a bridge in HomeKit.
// It's fine when the bridge has no explicit services.
//
// All accessory specific data (crypto keys, ids) is stored in a folder named after the first accessory.
// So changing the order of the accessories or renaming the first accessory makes the stored
// data inaccessible to the tranport. In this case new crypto keys are created and the accessory
// appears as a new one to clients.
func NewHAPTransport(config *HAPConfig, a *accessory.Accessory, as ...*accessory.Accessory) (Transport, error) {
	// Find transport name which is visible in mDNS
	name := a.Name()
	if len(name) == 0 || len(config.PIN) == 0 {
		log.Fatal("Invalid empty name for first accessory or invalid pin")
	}

	hapPin, err := NewPin(config.PIN)
	if err != nil {
		return nil, err
	}
        location := name
        if len(config.StoragePath) > 0 {
                location = path.Join(config.StoragePath, name)
        }
	storage, err := util.NewFileStorage(location)
	if err != nil {
		return nil, err
	}

	// Find transport uuid which appears as "id" txt record in mDNS and
	// must be unique and stay the same over time
	uuid := transportUUIDInStorage(storage)
	database := db.NewDatabaseWithStorage(storage)
	device, err := netio.NewSecuredDevice(uuid, hapPin, database)

	t := &ipTransport{
		database:  database,
		name:      name,
		device:    device,
		container: container.NewContainer(),
		mutex:     &sync.Mutex{},
		context:   netio.NewContextForSecuredDevice(device),
		emitter:   event.NewEmitter(),
                config:    config,
	}

	t.addAccessory(a)
	for _, a := range as {
		t.addAccessory(a)
	}

	t.emitter.AddListener(t)

	return t, err
}

func (t *ipTransport) Start() {
	s := server.NewServer(t.context, t.database, t.container, t.device, t.mutex, t.emitter)
	t.server = s
	port := to.Int64(s.Port())

	mdns := NewMDNSService(t.name, t.device.Name(), int(port))
	t.mdns = mdns

	if t.isPaired() {
		// Paired accessories must not be reachable for other clients since iOS 9
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
