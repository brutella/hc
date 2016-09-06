package hc

import (
	"errors"
	"bytes"
	"os"
	"io/ioutil"
	"net"
	"sync"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/event"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/hap/http"
	"github.com/brutella/hc/util"
	"github.com/brutella/log"
	"github.com/gosexy/to"
)

type ipTransportAlt struct {
	config  *Config
	context hap.Context
	server  http.Server
	mutex   *sync.Mutex
	mdns    *MDNSService

	storage  util.Storage
	database db.Database

	device    hap.SecuredDevice
	container *accessory.Container

	// Used to communicate between different parts of the program (e.g. successful pairing with HomeKit)
	emitter event.Emitter
}

// NewIPTransportAlt creates a transport to provide accessories over IP.
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
func NewIPTransportAlt(config Config, nodename string, a *accessory.Accessory) (Transport, error) {
	// Find transport name which is visible in mDNS
	name := a.Info.Name.GetValue()
	if len(name) == 0 {
		return nil, errors.New("invalid empty name for first accessory")
	}

	cfg := defaultConfig(name)
	cfg.merge(config)
	cfg.StoragePath = cfg.StoragePath  + "/" + nodename

	// Create a directory for this node
    if _, err := os.Stat(cfg.StoragePath); os.IsNotExist(err) {
		if err = os.MkdirAll(cfg.StoragePath, 0777); err != nil {
			return nil, err
		}
    }

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

	t := &ipTransportAlt{
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

	return t, err
}

// Add an accessory to the transport - must be done before the transport is started
func (t *ipTransportAlt) AddAccessory(a *accessory.Accessory) error {
	t.addAccessory(a)
	
	return nil
}

// Returns accessory that have the specified id
func (t *ipTransportAlt) GetAccessoryByID(aid int64) (a *accessory.Accessory) {
	return t.container.GetAccessoryByID(aid)
}

// Get the container holding all the accessories
func (t *ipTransportAlt) GetContainer() (container *accessory.Container) {
	return t.container
}

func (t *ipTransportAlt) Start() {
	// Users can only pair discoverable accessories
	if t.isPaired() {
		t.config.discoverable = false
	}

	t.config.categoryId = int(t.container.AccessoryType())
	t.config.updateConfigHash(t.container.ContentHash())
	t.config.save(t.storage)

	// Listen for events to update mDNS txt records
	t.emitter.AddListener(t)

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
	log.Printf("[INFO] Accessory IP is %s:%d", t.config.IP, t.config.servePort)

	// Listen until server.Stop() is called
	s.ListenAndServe()
}

// Stop stops the ip transport by unpublishing the mDNS service.
func (t *ipTransportAlt) Stop() {
	if t.mdns != nil {
		log.Printf("[INFO] Stop mDNS")
		t.mdns.Stop()
	}

	if t.server != nil {
		log.Printf("[INFO] Stop server. Accessory IP is %s:%d", t.config.IP, t.config.servePort)
		t.server.Stop()
	}
}

// Delete data related to this transport from disk
func (t *ipTransportAlt) RemoveFromDisk() error {
	return nil
}

// isPaired returns true when the transport is already paired
func (t *ipTransportAlt) isPaired() bool {

	// If more than one entity is stored in the database, we are paired with a device.
	// The transport itself is a device and is stored in the database, therefore
	// we have to check for more than one entity.
	if es, err := t.database.Entities(); err == nil && len(es) > 1 {
		return true
	}

	return false
}

func (t *ipTransportAlt) updateMDNSReachability() {
	if mdns := t.mdns; mdns != nil {
		t.config.discoverable = t.isPaired() == false
		mdns.Update()
	}
}

func (t *ipTransportAlt) addAccessory(a *accessory.Accessory) {
	t.container.AddAccessory(a)
	
	for _, s := range a.Services {
		for _, c := range s.Characteristics {
			// When a characteristic value changes and events are enabled for this characteristic
			// all listeners are notified. Since we don't track which client is interested in
			// which characteristic change event, we send them to all active connections.
			onConnChange := func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
				log.Printf("[VERB] onConnChange")
				if c.Events == true {
					t.notifyListener(a, c, conn)
				}
			}
			c.OnValueUpdateFromConn(onConnChange)

			onChange := func(c *characteristic.Characteristic, new, old interface{}) {
				log.Printf("[VERB] onChange")
				if c.Events == true {
					log.Printf("[VERB] onChange; c.Events == true")
					t.notifyListener(a, c, nil)
				}
			}
			c.OnValueUpdate(onChange)
		}
	}
}

func (t *ipTransportAlt) notifyListener(a *accessory.Accessory, c *characteristic.Characteristic, except net.Conn) {
	conns := t.context.ActiveConnections()
	for _, conn := range conns {
		if conn == except {
			continue
		}
		resp, err := hap.NewNotification(a, c)
		if err != nil {
			log.Fatal(err)
		}

		// Write response into buffer to replace HTTP protocol
		// specifier with EVENT as required by HAP
		var buffer = new(bytes.Buffer)
		resp.Write(buffer)
		bytes, err := ioutil.ReadAll(buffer)
		bytes = hap.FixProtocolSpecifier(bytes)
		log.Printf("[VERB] %s <- %s", conn.RemoteAddr(), string(bytes))
		conn.Write(bytes)
	}
}

// Handles event which are sent when pairing with a device is added or removed
func (t *ipTransportAlt) Handle(ev interface{}) {
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
