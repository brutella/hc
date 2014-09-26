package main

import(
    "fmt"
    "strconv"
    "net/http"
    
    "github.com/brutella/hap"
    "github.com/brutella/hap/pair"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/server"
)

var API_PORT int = 1234

// Announce service _hap._tcp via dns-sd
// dns-sd -R Accessory\ A _hap local 1234 pv=1.0 id=b1:42:90:21:73:9d c#=1 s#=1 sf=1 ff=0 md=HAP-Model-Name
func main() {
    bridge_info := model.NewAccessoryInfoService("123-456-789", "Rev1", "Matthias Hochgatterer", "My Bridge")        
    bridge_accessory := model.NewAccessory()
    bridge_accessory.AddService(bridge_info.Service)
    
    thermostat_info := model.NewAccessoryInfoService("001", "Model1a", "Matthias Hochgatterer", "Thermostat")
    thermostat_service := model.NewThermostatService("Schlafzimmer", 25, -20, 200, 1.0)
    thermostat_accessory := model.NewAccessory()
    thermostat_accessory.AddService(thermostat_info.Service)
    thermostat_accessory.AddService(thermostat_service.Service)
    
    m := model.NewModel()
    m.AddAccessory(bridge_accessory)
    m.AddAccessory(thermostat_accessory)
    
    model_controller := model.NewModelController(m)
    
    bridge, _   := hap.NewBridge("b1:42:90:21:73:9d", "001-02-003")
    storage, _  := hap.NewFileStorage("./ltpks")
    context     := hap.NewContext(storage)
    setup, _    := pair.NewSetupServerController(context, bridge)
    verify, _   := pair.NewVerifyServerController(context, bridge)
    
    mux :=  http.NewServeMux()
    
    setup_handler := server.NewPairSetupHandler(setup)
    mux.Handle("/pair-setup", setup_handler)
    
    verify_handler := server.NewPairVerifyHandler(verify, context)
    mux.Handle("/pair-verify", verify_handler)
    
    accessories_handler := server.NewAccessoriesHandler(model_controller, context)
    mux.Handle("/accessories", accessories_handler)
    
    addr := ":" + strconv.Itoa(API_PORT)
    fmt.Println("Running at", addr)
    err := server.ListenAndServe(addr, mux, context)
    fmt.Println(err)
}