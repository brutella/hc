package handler

import(
    "github.com/brutella/hap/pair"
    "github.com/brutella/hap/server"
        
    "io/ioutil"
    "net/http"
    "fmt"
)

type PairSetupHandler struct {
    http.Handler
    
    controller *pair.SetupServerController
}

func NewPairSetupHandler(c *pair.SetupServerController) *PairSetupHandler {
    handler := PairSetupHandler{
                controller: c,
            }
    
    return &handler
}

func (handler *PairSetupHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("POST /pair-setup")
    response.Header().Set("Content-Type", server.HTTPContentTypePairingTLV8)
    
    res, err := pair.HandleReaderForHandler(request.Body, handler.controller)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        bytes, _ := ioutil.ReadAll(res)
        response.Write(bytes)
    }
}