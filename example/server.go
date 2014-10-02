package main

import(
    "fmt"
    
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/common"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
    "github.com/brutella/hap/model/service"
    "github.com/brutella/hap/server"
    "github.com/brutella/hap/netio"
)

var API_PORT int = 1237

func main() {
    storage, _  := common.NewFileStorage("./data")
    database    := db.NewDatabaseWithStorage(storage)
    config      := netio.NewBridgeInfo("GoBridge", "001-02-003", "Matthias H.", storage)
    bridge, _   := netio.NewBridge(config)
    context     := netio.NewContextForBridge(bridge)
    
    fmt.Println("Run bridge")
    fmt.Println("            Name:", config.Name)
    fmt.Println("        Password:", config.Password)
    fmt.Println("   Serial Number:", config.SerialNumber)
    fmt.Println("              ID:", config.Id)
    
    bridge_info := service.NewAccessoryInfo(config.Name, config.SerialNumber, config.Manufacturer, "Bridge")
    bridge_accessory := accessory.NewAccessory()
    bridge_accessory.AddService(bridge_info.Service)
    
    thermostat_name := "Thermostat"
    thermostat_serial := common.GetSerialNumberForAccessoryName(thermostat_name, storage)
    thermostat_info := service.NewAccessoryInfo(thermostat_name, thermostat_serial, "Matthias H.", "Model1a")        
    thermostat_service := service.NewThermostat("Temperature", 20.9, 0.0, 100.0,  0.1)
    thermostat_accessory := accessory.NewAccessory()
    thermostat_accessory.AddService(thermostat_info.Service)
    thermostat_accessory.AddService(thermostat_service.Service)
    
    switch_name := "Smart Switch"
    switch_serial := common.GetSerialNumberForAccessoryName(switch_name, storage)
    switch_info := service.NewAccessoryInfo(switch_name, switch_serial, "Matthias H.", "Model1a")        
    switch_service := service.NewSwitch("Switch", true)
    switch_service.OnStateChanged(func(on bool){
        if on == true {
            fmt.Println("Switch is on")
        } else {
            fmt.Println("Switch is off")
        }
    })
    
    switch_accessory := accessory.NewAccessory()
    switch_accessory.AddService(switch_info.Service)
    switch_accessory.AddService(switch_service.Service)
    
    m := model.NewModel()
    m.AddAccessory(bridge_accessory)
    m.AddAccessory(thermostat_accessory)
    m.AddAccessory(switch_accessory)
    
    s := server.NewServer(context, database, m, bridge, API_PORT)
    
    fmt.Println("Publish service")
    fmt.Println("    ", s.DNSSDCommand())
    err := s.ListenAndServe()
    fmt.Println(err)
}