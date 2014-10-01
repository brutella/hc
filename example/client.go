package main

import(
    "fmt"
    "net/http"
    "github.com/brutella/hap"
    "github.com/brutella/hap/pair"
    "io"
    "os"
)

func sendTLV8(b io.Reader) (io.Reader, error){
    resp, err := http.Post("http://127.0.0.1:55036/pair-setup", hap.HTTPContentTypePairingTLV8, b)
    return resp.Body, err
}

func main() {    
    storage, err := hap.NewFileStorage(os.TempDir())
    context := hap.NewContext(storage)
    info := hap.NewBridgeInfo("Test Bridge", "719-47-107", "Matthias H.", storage)
    info.Id = "42:cd:02:57:0d:40"
    bridge, err := hap.NewBridge(info)
    
    client := pair.NewSetupClientController(context, bridge, "HomeKit Client")
    pairStartRequest := client.InitialPairingRequest()
    
    pairStartRespond, err := sendTLV8(pairStartRequest)
    if err != nil {
        fmt.Println(err)
    }
    
    // 2) S -> C
    pairVerifyRequest, err := client.HandleReader(pairStartRespond)
    if err != nil {
        fmt.Println(err)
    }
    
    // 3) C -> S
    pairVerifyRespond, err := sendTLV8(pairVerifyRequest)
    if err != nil {
        fmt.Println(err)
    }
    
    // 4) S -> C
    pairKeyRequest, err := client.HandleReader(pairVerifyRespond)
    if err != nil {
        fmt.Println(err)
    }
    
    // 5) C -> S
    pairKeyRespond, err := sendTLV8(pairKeyRequest)
    if err != nil {
        fmt.Println(err)
    }
    
    // 6) S -> C
    request, err := client.HandleReader(pairKeyRespond)
    if err != nil {
        fmt.Println(err)
    }
    
    if request != nil {
        fmt.Println(request)
    }
    
    fmt.Println("*** Pairing done ***")
    
    name := "UnitTest"
    verify := pair.NewVerifyClientController(context, bridge, name)
    
    verifyStartRequest := verify.InitialKeyVerifyRequest()
    // 1) C -> S
    verifyStartRespond, err := sendTLV8(verifyStartRequest)
    if err != nil {
        fmt.Println(err)
    }

    // 2) S -> C
    verifyFinishRequest, err := verify.Handle(verifyStartRespond)
    if err != nil {
        fmt.Println(err)
    }
    
    // 3) C -> S
    verifyFinishRespond, err := sendTLV8(verifyFinishRequest)
    if err != nil {
        fmt.Println(err)
    }
    
    // 4) S -> C 
    last_request, err := verify.Handle(verifyFinishRespond)
    if err != nil {
        fmt.Println(err)
    }
    
    if last_request != nil {
        fmt.Println(last_request)
    }
    
    fmt.Println("*** Key Verification done ***")
}