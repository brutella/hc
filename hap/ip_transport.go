package hap

import (
	"bytes"
	"io/ioutil"
	"log"
	"sync"

	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/container"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/netio/event"
	"github.com/brutella/hc/server"
	"github.com/gosexy/to"
)

// OnStopFunc is the function which is invoked when the transport stops.
type OnStopFunc func()

type ipTransport struct {
	context netio.HAPContext
	server  server.Server
	mutex   *sync.Mutex
	mdns    *MDNSService

	storage  common.Storage
	database db.Database

	stopFunc OnStopFunc

	name      string
	device    netio.SecuredDevice
	container *container.Container
}

func NewIPTransport(password string, as ...*accessory.Accessory) (Transport, error) {
	if len(as) == 0 {
		log.Fatalf("Invalid number of acccessories %d", len(as))
	}

	// Find a name for the transport automatically
	var name string
	for _, a := range as {
		if len(name) == 0 {
			name = a.Name()
		}
	}

	if len(name) == 0 {
		log.Fatal("IP Transport has no name")
	}

	storage, err := common.NewFileStorage(name)
	if err != nil {
		return nil, err
	}
    
    uuid := transportUUIDInStorage(storage)
	database := db.NewDatabaseWithStorage(storage)
	hapPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	device, err := netio.NewSecuredDevice(uuid, hapPassword, database)

	t := &ipTransport{
		database:  database,
		name:      name,
		device:    device,
		container: container.NewContainer(),
		mutex:     &sync.Mutex{},
		context:   netio.NewContextForSecuredDevice(device),
	}

	for _, a := range as {
		t.addAccessory(a)
	}

	return t, err
}

func (t *ipTransport) Start() {
	s := server.NewServer(t.context, t.database, t.container, t.device, t.mutex)
	port := to.Int64(s.Port())

	mdns := NewMDNSService(t.name, t.device.PairUsername(), int(port))
	t.mdns = mdns

	dns := t.database.DNSWithName(t.name)
	if dns == nil {
		dns = db.NewDNS(t.name, 1, 1)
		t.database.SaveDNS(dns)
	}
	mdns.Publish()

	s.OnStop(func() {
		t.Stop()
	})
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (t *ipTransport) OnStop(fn OnStopFunc) {
	t.stopFunc = fn
}

// Stop stops the ip transport by unpublishing the mDNS service.
func (t *ipTransport) Stop() {
	t.mdns.Stop()
	if t.stopFunc != nil {
		t.stopFunc()
	}
}

func (t *ipTransport) addAccessory(a *accessory.Accessory) {
	t.container.AddAccessory(a)

	for _, s := range a.Services {
		for _, c := range s.Characteristics {
			onChange := func(c *characteristic.Characteristic, new, old interface{}) {
				// (brutella) It's not clear yet when the state (s#) field in the TXT records
				// is updated. Sometimes it's increment when a client changes a value.

				// When a characteristic value changes and events are enabled for this characteristic
				// all listeners are notified. Since we don't track which client is interested in
				// which characteristic change event, we send them to all active connections.
				if c.Events == true {
					t.notifyListener(a, c)
				}
			}

			c.OnLocalChange(onChange)
			c.OnRemoteChange(onChange)
		}
	}
}

func (t *ipTransport) notifyListener(a *accessory.Accessory, c *characteristic.Characteristic) {
	conns := t.context.ActiveConnections()
	for _, con := range conns {
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
		log.Println("[VERB] <- ", string(bytes))
		con.Write(bytes)
	}
}

func transportUUIDInStorage(storage common.Storage) string {
	b_uuid, err := storage.Get("uuid")
	if len(b_uuid) == 0 || err != nil {
		str := common.RandomHexString()
		b_uuid = []byte(netio.MAC48Address(str))
		err := storage.Set("uuid", b_uuid)
		if err != nil {
			log.Fatal(err)
		}
	}
	return string(b_uuid)
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
