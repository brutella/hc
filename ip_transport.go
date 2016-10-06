package hc

import (
	"bytes"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/event"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/hap/http"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/util"
	"github.com/gosexy/to"
)

type ipTransport struct {
	config    *Config
	context   hap.Context
	server    http.Server
	keepAlive *hap.KeepAlive
	mutex     *sync.Mutex
	mdns      *MDNSService

	storage  util.Storage
	database db.Database

	device    hap.SecuredDevice
	container *accessory.Container

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
	name := a.Info.Name.GetValue()
	if len(name) == 0 {
		log.Info.Panic("Invalid empty name for first accessory")
	}

	cfg := defaultConfig(name)
	cfg.merge(config)

	storage, err := util.NewFileStorage(cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	database := db.NewDatabaseWithStorage(storage)

	hap_pin, err := NewPin(cfg.Pin)
	if err != nil {
		return nil, err
	}

	cfg.load(storage)

	device, err := hap.NewSecuredDevice(cfg.id, hap_pin, database)

	t := &ipTransport{
		storage:   storage,
		database:  database,
		device:    device,
		config:    cfg,
		container: accessory.NewContainer(),
		mutex:     &sync.Mutex{},
		context:   hap.NewContextForSecuredDevice(device),
		emitter:   event.NewEmitter(),
	}

	t.addAccessory(a)
	for _, a := range as {
		t.addAccessory(a)
	}

	// Users can only pair discoverable accessories
	if t.isPaired() {
		cfg.discoverable = false
	}

	cfg.categoryId = int(t.container.AccessoryType())
	cfg.updateConfigHash(t.container.ContentHash())
	cfg.save(storage)

	// Listen for events to update mDNS txt records
	t.emitter.AddListener(t)

	return t, err
}

func (t *ipTransport) Start() {

	// Create server which handles incoming tcp connections
	config := http.Config{
		Port:      t.config.Port,
		Context:   t.context,
		Database:  t.database,
		Container: t.container,
		Device:    t.device,
		Mutex:     t.mutex,
		Emitter:   t.emitter,
	}

	s := http.NewServer(config)
	t.server = s

	// Publish server port which might be different then `t.config.Port`
	t.config.servePort = int(to.Int64(s.Port()))

	mdns := NewMDNSService(t.config)
	t.mdns = mdns

	mdns.Publish()

	// Publish accessory ip
	log.Info.Println("Accessory IP is", t.config.IP)

	// Send keep alive notifications to all connected clients every 10 minutes
	t.keepAlive = hap.NewKeepAlive(10*time.Minute, t.context)
	go t.keepAlive.Start()

	// Listen until server.Stop() is called
	s.ListenAndServe()
}

// Stop stops the ip transport by unpublishing the mDNS service.
func (t *ipTransport) Stop() {

	if t.keepAlive != nil {
		t.keepAlive.Stop()
	}

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
		t.config.discoverable = t.isPaired() == false
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
			c.OnValueUpdateFromConn(onConnChange)

			onChange := func(c *characteristic.Characteristic, new, old interface{}) {
				if c.Events == true {
					t.notifyListener(a, c, nil)
				}
			}
			c.OnValueUpdate(onChange)
		}
	}
}

func (t *ipTransport) notifyListener(a *accessory.Accessory, c *characteristic.Characteristic, except net.Conn) {
	conns := t.context.ActiveConnections()
	for _, conn := range conns {
		if conn == except {
			continue
		}
		resp, err := hap.NewCharacteristicNotification(a, c)
		if err != nil {
			log.Info.Panic(err)
		}

		// Write response into buffer to replace HTTP protocol
		// specifier with EVENT as required by HAP
		var buffer = new(bytes.Buffer)
		resp.Write(buffer)
		bytes, err := ioutil.ReadAll(buffer)
		bytes = hap.FixProtocolSpecifier(bytes)
		log.Debug.Printf("%s <- %s", conn.RemoteAddr(), string(bytes))
		conn.Write(bytes)
	}
}

// Handles event which are sent when pairing with a device is added or removed
func (t *ipTransport) Handle(ev interface{}) {
	switch ev.(type) {
	case event.DevicePaired:
		log.Debug.Printf("Event: paired with device")
		t.updateMDNSReachability()
	case event.DeviceUnpaired:
		log.Debug.Printf("Event: unpaired with device")
		t.updateMDNSReachability()
	default:
		break
	}
}
