package server

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap"
    "github.com/brutella/hap/pair"
)

type PairingHandler struct {
    http.Handler
    controller *pair.PairingController
    context *hap.Context
}

func NewPairingHandler(controller *pair.PairingController, context *hap.Context) *PairingHandler {
    handler := PairingHandler{
                controller: controller,
                context: context,
            }
    
    return &handler
}

func (handler *PairingHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("POST /pairings")
    
    _, err := pair.HandleReaderForHandler(request.Body, handler.controller)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        response.WriteHeader(http.StatusNoContent)
    }
}