package main

import(
    "fmt"
    "strconv"
    "net/http"
    
    "github.com/brutella/hap"
    "github.com/brutella/hap/pair"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/example/server/src"
)

var API_PORT int = 1234

// Announce service _hap._tcp via dns-sd
// dns-sd -R Accessory\ A _hap local 1234 pv=1.0 id=b1:42:90:21:73:9d c#=1 s#=1 sf=1 ff=0 md=HAP-Model-Name
func main() {
    info_service := model.NewAccessoryInfoService("123-456-789", "Rev1", "Matthias H.", "My Bridge")
    accessory := model.NewAccessory()
    accessory.AddService(info_service.Service)
    m := model.NewModel()
    m.AddAccessory(accessory)
    model_controller := model.NewModelController(m)
    
    bridge, _ := hap.NewBridge("b1:42:90:21:73:9d", "001-02-003")
    storage, _ := hap.NewFileStorage("./ltpks")
    context := hap.NewContext(storage)
    setup, _ := pair.NewSetupServerController(context, bridge)
    verify, _ := pair.NewVerifyServerController(context, bridge)
    
    mux :=  http.NewServeMux()
    
    setup_handler := hapserver.NewPairSetupHandler(setup)
    mux.Handle("/pair-setup", setup_handler)
    
    verify_handler := hapserver.NewPairVerifyHandler(verify)
    mux.Handle("/pair-verify", verify_handler)
    
    accessories_handler := hapserver.NewAccessoriesHandler(model_controller, context)
    mux.Handle("/accessories", accessories_handler)
        
    fmt.Println("Running at :" + strconv.Itoa(API_PORT)) 
    http.ListenAndServe(":" + strconv.Itoa(API_PORT), mux)   
}