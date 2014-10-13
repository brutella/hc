package app

import (
    "log"
    "errors"
    
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/common"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/container"
    "github.com/brutella/hap/model/accessory"
    "github.com/brutella/hap/server"
    "github.com/brutella/hap/netio"
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
        BridgePassword: "00102003",
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
        Serial: bridge_config.SerialNumber,
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
}

func (app *App) Run() {
    s := server.NewServer(app.context, app.Database, app.container, app.bridge)
    
    s.OnExit(func() {
        if app.exitFunc != nil {
            app.exitFunc()
        }
    })
    err := s.ListenAndServe()
    log.Fatal(err)
}

func (app *App) OnExit(fn AppExitFunc) {
    app.exitFunc = fn
}
