package hap

import (
	"bytes"
	"io/ioutil"
	"net"
	"sync"

	"github.com/brutella/hc/db"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/container"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/netio/event"
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
}

// NewIPTransport creates a transport to provide accessories over IP.
// The pairing is secured using a 8-numbers password.
// If more than one accessory is provided, the first becomes a bridge in HomeKit.
// It's fine when the bridge has no explicit services.
//
// All accessory specific data (crypto keys, ids) is stored in a folder named after the first accessory.
// So changing the order of the accessories or renaming the first accessory makes the stored
// data inaccessible to the tranport. In this case new crypto keys are created and the accessory
// appears as a new one to clients.
func NewIPTransport(password string, a *accessory.Accessory, as ...*accessory.Accessory) (Transport, error) {
	// Find transport name which is visible in mDNS
	name := a.Name()
	if len(name) == 0 {
		log.Fatal("Invalid empty name for first accessory")
	}

	hapPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	storage, err := util.NewFileStorage(name)
	if err != nil {
		return nil, err
	}

	// Find transport uuid which appears as "id" txt record in mDNS and
	// must be unique and stay the same over time
	uuid := transportUUIDInStorage(storage)
	database := db.NewDatabaseWithStorage(storage)
	device, err := netio.NewSecuredDevice(uuid, hapPassword, database)

	t := &ipTransport{
		database:  database,
		name:      name,
		device:    device,
		container: container.NewContainer(),
		mutex:     &sync.Mutex{},
		context:   netio.NewContextForSecuredDevice(device),
	}

	t.addAccessory(a)
	for _, a := range as {
		t.addAccessory(a)
	}

	return t, err
}

func (t *ipTransport) Start() {
	s := server.NewServer(t.context, t.database, t.container, t.device, t.mutex)
	t.server = s
	port := to.Int64(s.Port())

	mdns := NewMDNSService(t.name, t.device.Name(), int(port))
	t.mdns = mdns

	dns := t.database.DNSWithName(t.name)
	if dns == nil {
		dns = db.NewDNS(t.name, 1, 1)
		t.database.SaveDNS(dns)
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
		resp, err := event.New(a, c)
		if err != nil {
			log.Fatal(err)
		}

		// Write response into buffer to replace HTTP protocol
		// specifier with EVENT as required by HAP
		var buffer = new(bytes.Buffer)
		resp.Write(buffer)
		bytes, err := ioutil.ReadAll(buffer)
		bytes = event.FixProtocolSpecifier(bytes)
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

// updateConfiguration increases the configuration value by 1 and re-announces the new TXT records.
// This method is currently not used.
func (t *ipTransport) updateConfiguration() {
	dns := t.database.DNSWithName(t.name)
	if dns != nil {
		dns.SetConfiguration(dns.Configuration() + 1)
		t.database.SaveDNS(dns)
		if t.mdns != nil {
			t.mdns.Update()
		}
	}
}
