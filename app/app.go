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
    
    bridge  *netio.Bridge
    container   *container.Container
    exitFunc AppExitFunc
    mdns *Service
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
            c.OnLocalChange(func(c *characteristic.Characteristic, oldValue interface{}) {
                // (brutella) It's not clear yet when the state (s#) field in the TXT records
                // is updated. Sometimes it's increment when a client changes a value.
                // if app.mdns != nil {
                //     log.Println("[VERB] Update TXT records")
                //     app.mdns.state += 1
                //     app.mdns.Update()
                // }
                
                if c.Events == true {
                    app.NotifyListener(a, c)
                }
            })
        }
    }
}

func (app *App) RemoveAccessory(a *accessory.Accessory) {
    app.container.RemoveAccessory(a)
}

func (app *App) Run() {
    s := server.NewServer(app.context, app.Database, app.container, app.bridge)
    s.OnStop(func() {
        // Stop mDNS
        if app.mdns != nil {
            log.Println("[INFO] Stop mdns")
            app.mdns.Stop()
        }
        
        if app.exitFunc != nil {
            app.exitFunc()
        }
    })
    
    go func() {
        time.Sleep(1 * time.Second)
        app.PublishServer(s)
    }()
    
    err := s.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}

func (app *App) OnExit(fn AppExitFunc) {
    app.exitFunc = fn
}

func (app *App) PublishServer(server server.Server) {
    port := to.Int64(server.Port())
    // TODO Store state and configuration on disk
    mdns := NewService(app.bridge.Name(), app.bridge.Id(), int(port))
    err := mdns.Publish()
    if err != nil {
        log.Fatal("Could not publish server", err)
    }
    
    app.mdns = mdns
}

func (app *App) NotifyListener(a *accessory.Accessory, c *characteristic.Characteristic) {
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