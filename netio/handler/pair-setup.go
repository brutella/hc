package handler

import(
    "github.com/brutella/hap/netio/pair"
    "github.com/brutella/hap/netio"
        
    "io/ioutil"
    "net/http"
    "fmt"
)

type PairSetupHandler struct {
    http.Handler
    
    controller *pair.SetupServerController
    context netio.Context
}

func NewPairSetupHandler(c *pair.SetupServerController, context netio.Context) *PairSetupHandler {
    handler := PairSetupHandler{
                controller: c,
                context: context,
            }
    
    return &handler
}

func (handler *PairSetupHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("POST /pair-setup")
    response.Header().Set("Content-Type", netio.HTTPContentTypePairingTLV8)
    
    res, err := pair.HandleReaderForHandler(request.Body, handler.controller)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        bytes, _ := ioutil.ReadAll(res)
        response.Write(bytes)
    }
}