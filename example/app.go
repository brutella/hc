package main

import (
    "log"
    "time"
    
    "github.com/brutella/hap/app"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
)

func main() {
    conf := app.NewConfig()
    conf.DatabaseDir = "./data"
    
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
    
    go func() {
        timer := time.NewTimer(2 * time.Second)
        for {
            <- timer.C
            log.Println("Update switch")
            sw.SetOn(sw.IsOn() == false)
            timer.Reset(2 * time.Second)
        }
    }()
    
    app.Run()
}