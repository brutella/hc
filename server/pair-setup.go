package server

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap/pair"
    "github.com/brutella/hap"
    "io/ioutil"
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
    response.Header().Set("Content-Type", hap.HTTPContentTypePairingTLV8)
    
    res, err := handler.controller.Handle(request.Body)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        bytes, _ := ioutil.ReadAll(res)
        response.Write(bytes)
    }
}