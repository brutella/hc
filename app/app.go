package app

import (
    "errors"
    "io/ioutil"
    "bytes"
    "sync"
    
    "github.com/brutella/log"
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/common"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/container"
    "github.com/brutella/hap/model/accessory"
    "github.com/brutella/hap/model/characteristic"
    "github.com/brutella/hap/server"
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/netio/event"
    "github.com/gosexy/to"
)

type Config struct {
    BridgeName string
    BridgePassword string
    BridgeManufacturer string
    DatabaseDir string
}

func NewConfig() Config {
    return Config{
        BridgeName: "GoBridge",
        BridgePassword: "001-02-003",
        BridgeManufacturer: "brutella",
    }
}

type AppExitFunc func()
type App struct {    
    context netio.HAPContext
    
    Database db.Database
    Storage  common.Storage
    
    bridge    *netio.Bridge
    server    server.Server
    
    mutex     *sync.Mutex
    mdns      *Service
    container *container.Container
    
    exitFunc AppExitFunc
    batchUpdate bool
}

func NewApp(conf Config) (*App, error) {
    if len(conf.DatabaseDir) == 0 {
        return nil, errors.New("Database directory not specified")
    }
    
    storage, err  := common.NewFileStorage(conf.DatabaseDir)
    if err != nil {
        return nil, err
    }
    
    database      := db.NewDatabaseWithStorage(storage)
    bridge_config := netio.NewBridgeInfo(conf.BridgeName, conf.BridgePassword, conf.BridgeManufacturer, storage)
    bridge, err   := netio.NewBridge(bridge_config)
    if err != nil {
        return nil, err
    }
    
    info := model.Info{
        Name: bridge_config.Name,
        SerialNumber: bridge_config.SerialNumber,
        Manufacturer: bridge_config.Manufacturer,
        Model: "Bridge",
    }
    bridge_accessory := accessory.New(info)
    container := container.NewContainer()
    container.AddAccessory(bridge_accessory)
    
    app := App{
        context: netio.NewContextForBridge(bridge),
        bridge: bridge,
        Storage: storage,
        Database: database,
        container: container,
        mutex: &sync.Mutex{},
    }
    
    return &app, nil
}

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

func (app *App) RemoveAccessory(a *accessory.Accessory) {
    app.container.RemoveAccessory(a)
    if app.batchUpdate == false {
        app.updateConfiguration()
    }
}

func (app *App) PerformBatchUpdates(fn func()) {
    app.mutex.Lock()
    defer app.mutex.Unlock()
    app.batchUpdate = true
    fn()
    app.batchUpdate = false
}

// SetReachable update app's reachability status
// When a Bonjour service is running and reachable is false, the Bonjour service is stopped
// When no Bonjour service is running and reachable is true, the service is announed via Bonjour
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

func (app *App) IsReachable() bool {
    return app.mdns != nil && app.mdns.IsPublished()
}

// Run starts the server and publishes the service via Bonjour
func (app *App) Run() {
    app.RunAndPublish(true)
}

// RunAndPublish starts the server
// If publish is true, the Bonjour service is started automatically
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
        mdns.Update()
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

func (app *App) OnExit(fn AppExitFunc) {
    app.exitFunc = fn
}

func (app *App) Stop() {
    if app.mdns != nil {
        app.mdns.Stop()
    }
    if app.exitFunc != nil {
        app.exitFunc()
    }
}

func (app *App) closeAllConnections() {
    for _, c := range app.context.ActiveConnections() {
        c.Close()
    }
}

func (app *App) notifyListener(a *accessory.Accessory, c *characteristic.Characteristic) {
    conns := app.context.ActiveConnections()    
    for _, con := range conns {
        resp, err := event.NewNotification(a, c)
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

// updateConfiguration increases the configuration value by 1 
// re-announced new Text records
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