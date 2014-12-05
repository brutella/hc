package app

import (
    "errors"
    "time"
    "io/ioutil"
    "bytes"
    
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
    Storage common.Storage
    
    bridge    *netio.Bridge
    container *container.Container
    server    server.Server
    mdns      *Service
    
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
    
    database    := db.NewDatabaseWithStorage(storage)
    bridge_config := netio.NewBridgeInfo(conf.BridgeName, conf.BridgePassword, conf.BridgeManufacturer, storage)
    
    bridge, err   := netio.NewBridge(bridge_config)
    if err != nil {
        return nil, err
    }
    context     := netio.NewContextForBridge(bridge)
    
    info := model.Info{
        Name: bridge_config.Name,
        SerialNumber: bridge_config.SerialNumber,
        Manufacturer: bridge_config.Manufacturer,
        Model: "Bridge",
    }
    
    bridge_accessory := accessory.New(info)
    
    cont := container.NewContainer()
    cont.AddAccessory(bridge_accessory)
    
    app := App{
        context: context,
        bridge: bridge,
        Storage: storage,
        Database: database,
        container: cont,
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
    if app.batchUpdate == false {
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
    app.batchUpdate = true
    fn()
    app.updateConfiguration()
    app.batchUpdate = false
}

// SetReachable update app's reachability status
// When a Bonjour service is running and reachable is false, the Bonjour service is stopped
// When no Bonjour service is running and reachable is true, the service is announed via Bonjour
func (app *App) SetReachable(reachable bool) {
    if app.mdns == nil && reachable == true {
        app.publishService(app.server)
    } else if app.mdns != nil && reachable == false {
        app.closeAllConnections()
        app.stopService()
    }
}

// Run starts the server and publishes the service via Bonjour
func (app *App) Run() {
    app.RunAndPublish(true)
}

// RunAndPublish starts the server
// If publish is true, the Bonjour service is started automatically
func (app *App) RunAndPublish(publish bool) {
    s := server.NewServer(app.context, app.Database, app.container, app.bridge)
    s.OnStop(func() {
        app.Stop()
    })
    
    if publish == true {
        go func() {
            time.Sleep(1 * time.Second)
            app.publishService(s)
        }()
    }
    
    app.server = s
    err := s.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}

func (app *App) OnExit(fn AppExitFunc) {
    app.exitFunc = fn
}

func (app *App) Stop() {
    // Stop mDNS
    app.stopService()
    if app.exitFunc != nil {
        app.exitFunc()
    }
}

func (app *App) publishService(server server.Server) {
    port := to.Int64(server.Port())
    mdns := NewService(app.bridge.Name(), app.bridge.Id(), int(port))
    
    // c# and s# TXT records are stored in database
    // Set to 1 on first launch
    dns := app.Database.DnsWithName(app.bridge.Name())
    if dns == nil {
        dns = db.NewDns(app.bridge.Name(), 1, 1)
        app.Database.SaveDns(dns)
    }
    mdns.configuration = dns.Configuration()
    mdns.state = dns.State()
    
    err := mdns.Publish()
    if err != nil {
        log.Fatal("Could not publish server", err)
    }
    
    app.mdns = mdns
}

func (app *App) stopService() {
    if app.mdns != nil {
        log.Println("[INFO] Stopping mdns...")
        app.mdns.Stop()
        app.mdns = nil
    }
}

func (app *App) closeAllConnections() {
    for _, c := range app.context.ActiveConnection() {
        c.Close()
    }
}

func (app *App) notifyListener(a *accessory.Accessory, c *characteristic.Characteristic) {
    conns := app.context.ActiveConnection()    
    for _, con := range conns {
        resp, err := event.NewNotification(a, c)
        if err != nil {
            log.Fatal(err)
        }
        
        // Write response into buffer to replace HTTP protocol 
        // specifier with Event as required by HAP
        var buffer = new(bytes.Buffer)
        resp.Write(buffer)
        bytes, err := ioutil.ReadAll(buffer)
        bytes = event.FixProtocolSpecifier(bytes)
        log.Println("[VERB] <- ", string(bytes))
        
        // Write bytes to connection instead of using response object
        // resp.Write(con)
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
        app.updateDns()
    }
}

func (app *App) updateDns() {
    dns := app.Database.DnsWithName(app.bridge.Name())
    if app.mdns != nil && dns != nil {
        app.mdns.configuration = dns.Configuration()
        app.mdns.state = dns.State()
        app.mdns.Update()
    }
}