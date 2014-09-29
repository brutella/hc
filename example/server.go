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
// dns-sd -R GoBridge _hap local 1234 pv=1.0 id=a1:42:90:21:73:9d+1 c#=1 s#=1 sf=1 ff=0 md=GoBridge
//
// md – accessory model name (GoBridge)
// pv – protocol version, "1.0";
// id – identifier of device (accessory username, see below); (BBB = Beaglebone Black)
// c# – configuration number, incremented every time services for accessory are updated;
// s# – state number (seems to be unused, usually matches highest service number);
// sf – status flags (seems to be unused, always "1");
// ff – feature flags. If bit 0 is set, device is considered MFi-compliant, and additional verification is performed, otherwise a warning about non-compliant device is shown before PIN code request.

func main() {
    info := hap.NewBridgeInfo("GoBridge", "001-02-003", "0002", "Matthias H.")
    bridge_info := model.NewAccessoryInfoService(info.Name, info.SerialNumber, info.Manufacturer, "Bridge")
    bridge_accessory := model.NewAccessory()
    bridge_accessory.AddService(bridge_info.Service)
    
    thermostat_info := model.NewAccessoryInfoService("Raumtemperatur", "000-000-001", "Matthias H.", "Thermostat")        
    thermostat_service := model.NewThermostatService("Thermostat", 22.1, 0.0, 100.0,  0.1)
    thermostat_accessory := model.NewAccessory()
    thermostat_accessory.AddService(thermostat_info.Service)
    thermostat_accessory.AddService(thermostat_service.Service)
    
    m := model.NewModel()
    m.AddAccessory(bridge_accessory)
    m.AddAccessory(thermostat_accessory)
    
    model_controller := model.NewModelController(m)
    characteristics_controller := server.NewCharacteristicController(m)
    
    bridge, _   := hap.NewBridge(info)
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
    
    characteristics_handler := server.NewCharacteristicsHandler(characteristics_controller, context)
    mux.Handle("/characteristics", characteristics_handler)
    
    addr := ":" + strconv.Itoa(API_PORT)
    fmt.Println("Running at", addr)
    fmt.Println("Publish service")
    fmt.Printf("    dns-sd -R %s _hap local %s pv=1.0 id=%s c#=1 s#=1 sf=1 ff=0 md=%s\n", bridge.Name(), strconv.Itoa(API_PORT), info.Id, bridge.Name())
    err := server.ListenAndServe(addr, mux, context)
    fmt.Println(err)
}