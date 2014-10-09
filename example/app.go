package main

import (
    "log"
    
    "github.com/brutella/hap/app"
    "github.com/brutella/hap/server"
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
    app.Run()
}