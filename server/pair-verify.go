package server

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap/pair"
    "github.com/brutella/hap"
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
    response.Header().Set("Content-Type", hap.HTTPContentTypePairingTLV8)
    
    res, err := pair.HandleReaderForHandler(request.Body, handler.controller)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        bytes, _ := ioutil.ReadAll(res)
        response.Write(bytes)
    }
}