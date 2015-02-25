package hap

import (
	"bytes"
	"errors"
	"io/ioutil"
	"sync"

	"github.com/brutella/hap/common"
	"github.com/brutella/hap/db"
	"github.com/brutella/hap/model"
	"github.com/brutella/hap/model/accessory"
	"github.com/brutella/hap/model/characteristic"
	"github.com/brutella/hap/model/container"
	"github.com/brutella/hap/netio"
	"github.com/brutella/hap/netio/event"
	"github.com/brutella/hap/server"
	"github.com/brutella/log"
	"github.com/gosexy/to"
)

type AppExitFunc func()

// App encapsulates all components to create, publish and update accessories and
// make the available via mDNS.
type App struct {
	context netio.HAPContext

	Database db.Database
	Storage  common.Storage

	bridge *netio.Bridge
	server server.Server

	mutex     *sync.Mutex
	mdns      *Service
	container *container.Container

	exitFunc    AppExitFunc
	batchUpdate bool
}

// NewApp returns a app based on the configuration.
func NewApp(conf Config) (*App, error) {
	if len(conf.DatabaseDir) == 0 {
		return nil, errors.New("Database directory not specified")
	}

	storage, err := common.NewFileStorage(conf.DatabaseDir)
	if err != nil {
		return nil, err
	}

	database := db.NewDatabaseWithStorage(storage)
	bridge_config := netio.NewBridgeInfo(conf.BridgeName, conf.BridgePassword, conf.BridgeManufacturer, storage)
	bridge, err := netio.NewBridge(bridge_config)
	if err != nil {
		return nil, err
	}

	// Bridge appears in HomeKit and must provide the mandatory accessory info servie
	info := model.Info{
		Name:         bridge_config.Name,
		SerialNumber: bridge_config.SerialNumber,
		Manufacturer: bridge_config.Manufacturer,
		Model:        "Bridge",
	}
	bridge_accessory := accessory.New(info)
	container := container.NewContainer()
	container.AddAccessory(bridge_accessory)

	app := App{
		context:   netio.NewContextForBridge(bridge),
		bridge:    bridge,
		Storage:   storage,
		Database:  database,
		container: container,
		mutex:     &sync.Mutex{},
	}

	return &app, nil
}

// AddAccessory adds an accessory to the bridge and updates the mDNS configuration
// when no batch updates are currently active.
func (app *App) AddAccessory(a *accessory.Accessory) {
	app.container.AddAccessory(a)

	for _, s := range a.Services {
		for _, c := range s.Characteristics {
			onChange := func(c *characteristic.Characteristic, oldValue interface{}) {
				// (brutella) It's not clear yet when the state (s#) field in the TXT records
				// is updated. Sometimes it's increment when a client changes a value.
				// if app.mdns != nil {
				//     log.Println("[VERB] Update TXT records")
				//     app.mdns.state += 1
				//     app.mdns.Update()
				// }

				// When a characteristic value changes and events are enabled for this characteristic
				// all listeners are notified. Since we don't track which client is interested in
				// which characteristic change event, we send them to all active connections.
				if c.Events == true {
					app.notifyListener(a, c)
				}
			}

			c.OnLocalChange(onChange)
			c.OnRemoteChange(onChange)
		}
	}
	if app.batchUpdate == false && app.IsReachable() {
		app.updateConfiguration()
	}
}

// RemoveAccessory removes the accessory and updates the mDNS configuration
// when no batch updates are currently active.
func (app *App) RemoveAccessory(a *accessory.Accessory) {
	app.container.RemoveAccessory(a)
	if app.batchUpdate == false {
		app.updateConfiguration()
	}
}

// PerformBatchUpdates allows multiple accessory (adding, removing) and characteristic (value update)
// changes at once without triggering mDNS configuration updates after every change.
func (app *App) PerformBatchUpdates(fn func()) {
	app.mutex.Lock()
	defer app.mutex.Unlock()
	app.batchUpdate = true
	fn()
	app.batchUpdate = false
}

// SetReachable update the app's reachability status.
// When reachable is true, the app will be announed via mDNS and is then visible to HomeKit clients.
// When reachable is false, the app will be unannounced via mDNS and all connections get closed.
func (app *App) SetReachable(reachable bool) {
	if app.IsReachable() != reachable {
		if reachable == true {
			app.mdns.Publish()
		} else {
			app.closeAllConnections()
			app.mdns.Stop()
		}
	}
}

// IsReachable returns true when the app is reachable via mDNS, otherwise false.
func (app *App) IsReachable() bool {
	return app.mdns != nil && app.mdns.IsPublished()
}

// Run calls RunAndPublish(true)
func (app *App) Run() {
	app.RunAndPublish(true)
}

// RunAndPublish starts a TCP server which handles sockets on a random port. The method blocks until the server stopped.
// If publish is true, the mDNS service is started. Otherwise you have to call SetReachable(true) to make the app visible on the network.
//
// The app gracefully stops when the server received either an interrupt or kill signal.
func (app *App) RunAndPublish(publish bool) {
	s := server.NewServer(app.context, app.Database, app.container, app.bridge, app.mutex)
	port := to.Int64(s.Port())

	app.mutex.Lock()
	mdns := NewService(app.bridge.Name(), app.bridge.Id(), int(port))
	app.mdns = mdns
	app.mutex.Unlock()

	dns := app.Database.DnsWithName(app.bridge.Name())
	if dns == nil {
		dns = db.NewDns(app.bridge.Name(), 1, 1)
		app.Database.SaveDns(dns)
	}
	if publish {
		app.mutex.Lock()
		mdns.Publish()
		app.mutex.Unlock()
	}

	s.OnStop(func() {
		app.Stop()
	})
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// OnExit calls the argument function when the app is stopped and ready for termination.
func (app *App) OnExit(fn AppExitFunc) {
	app.exitFunc = fn
}

// Stop stops the app by unpublishing the mDNS service.
func (app *App) Stop() {
	if app.mdns != nil {
		app.mdns.Stop()
	}
	if app.exitFunc != nil {
		app.exitFunc()
	}
}

// closeAllConnections calls Close on all active connections
func (app *App) closeAllConnections() {
	for _, c := range app.context.ActiveConnections() {
		c.Close()
	}
}

// notifyListener sends an EVENT HTTP packet containing the characteristic value to all active connections
func (app *App) notifyListener(a *accessory.Accessory, c *characteristic.Characteristic) {
	conns := app.context.ActiveConnections()
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

// updateConfiguration increases the configuration value by 1 and re-announces the new TXT records
func (app *App) updateConfiguration() {
	dns := app.Database.DnsWithName(app.bridge.Name())
	if dns != nil {
		dns.SetConfiguration(dns.Configuration() + 1)
		app.Database.SaveDns(dns)
		if app.mdns != nil {
			app.mdns.Update()
		}
	}
}
