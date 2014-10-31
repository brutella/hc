package main

import (
    "log"
    
    "github.com/brutella/hap/app"
    "github.com/brutella/hap/server"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
)

func main() {
    conf := app.NewConfig()
    conf.DatabaseDir = "./data"
    conf.BridgeName = "TestBridge" // default "GoBridge"
    
    pwd, _ := server.NewPassword("11122333")
    conf.BridgePassword = pwd // default "001-02-003"
    conf.BridgeManufacturer = "Matthias Hochgatterer" // default "brutella"
    
    app, err := app.NewApp(conf)
    if err != nil {
        log.Fatal(err)
    }
    
    info := model.Info{
        Name: "My Switch",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Switchy",
    }
    
    sw := accessory.NewSwitch(info)
    sw.OnStateChanged(func(on bool) {
        if on == true {
            log.Println("Switch on")
        } else {
            log.Println("Switch off")
        }
    })
    
    app.AddAccessory(sw.Accessory)
    
    app.Run()
}