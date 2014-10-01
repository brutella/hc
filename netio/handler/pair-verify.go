package handler

import(
    "github.com/brutella/hap"
    "github.com/brutella/hap/pair"
    "github.com/brutella/hap/netio"
    
    "net/http"
    "fmt"
    "io/ioutil"
)

type PairVerifyHandler struct {
    http.Handler
    controller *pair.VerifyServerController
    context *hap.Context
}

func NewPairVerifyHandler(controller *pair.VerifyServerController, context *hap.Context) *PairVerifyHandler {
    handler := PairVerifyHandler{
                controller: controller,
                context: context,
            }
    
    return &handler
}

func (handler *PairVerifyHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("POST /pair-verify")
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