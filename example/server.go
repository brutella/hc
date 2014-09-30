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

var API_PORT int = 1235

// Announce service _hap._tcp via dns-sd
// dns-sd -R GoBridge _hap local 1234 pv=1.0 id=a1:42:90:21:73:9d c#=1 s#=1 sf=1 ff=0 md=GoBridge
//
// md – accessory model name (GoBridge)
// pv – protocol version, "1.0";
// id – identifier of device (accessory username, see below); (BBB = Beaglebone Black)
// c# – configuration number, incremented every time services for accessory are updated;
// s# – state number (seems to be unused, usually matches highest service number);
// sf – status flags (seems to be unused, always "1");
// ff – feature flags. If bit 0 is set, device is considered MFi-compliant, and additional verification is performed, otherwise a warning about non-compliant device is shown before PIN code request.

func main() {
    storage, _  := hap.NewFileStorage("./data")
    context     := hap.NewContext(storage)
    config      := hap.NewBridgeInfo("GoBridge", "001-02-003", "Matthias H.", storage)
    bridge, _   := hap.NewBridge(config)
    fmt.Println("Run bridge")
    fmt.Println("            Name:", config.Name)
    fmt.Println("        Password:", config.Password)
    fmt.Println("   Serial Number:", config.SerialNumber)
    fmt.Println("              ID:", config.Id)
    
    bridge_info := model.NewAccessoryInfoService(config.Name, config.SerialNumber, config.Manufacturer, "Bridge")
    bridge_accessory := model.NewAccessory()
    bridge_accessory.AddService(bridge_info.Service)
    
    thermostat_name := "Temperaturregler"
    thermostat_serial := hap.GetSerialNumberForAccessoryName(thermostat_name, storage)
    thermostat_info := model.NewAccessoryInfoService(thermostat_name, thermostat_serial, "Matthias H.", "Model1a")        
    thermostat_service := model.NewThermostatService("Temperatur", 20.9, 0.0, 100.0,  0.1)
    thermostat_accessory := model.NewAccessory()
    thermostat_accessory.AddService(thermostat_info.Service)
    thermostat_accessory.AddService(thermostat_service.Service)
    
    m := model.NewModel()
    m.AddAccessory(bridge_accessory)
    m.AddAccessory(thermostat_accessory)
    
    model_controller := model.NewModelController(m)
    characteristics_controller := server.NewCharacteristicController(m)
    
    setup, _    := pair.NewSetupServerController(context, bridge)
    verify, _   := pair.NewVerifyServerController(context, bridge)
    
    mux :=  http.NewServeMux()
    
    setup_handler := server.NewPairSetupHandler(setup)
    mux.Handle("/pair-setup", setup_handler)
    
    verify_handler := server.NewPairVerifyHandler(verify, context)
    mux.Handle("/pair-verify", verify_handler)
    
    accessories_handler := server.NewAccessoriesHandler(model_controller, context)
    mux.Handle("/accessories", accessories_handler)
    
    characteristics_handler := server.NewCharacteristicsHandler(characteristics_controller, context)
    mux.Handle("/characteristics", characteristics_handler)
    
    addr := ":" + strconv.Itoa(API_PORT)
    fmt.Println("Running at", addr)
    fmt.Println("Publish service")
    fmt.Printf("    dns-sd -P %s _hap local %s macbookpro.local 192.168.0.14 pv=1.0 id=%s c#=1 s#=1 sf=1 ff=0 md=%s\n", bridge.Name(), strconv.Itoa(API_PORT), bridge.Id(), bridge.Name())
    err := server.ListenAndServe(addr, mux, context)
    fmt.Println(err)
}