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
    
    tlv8, err := pair.NewTLV8ContainerFromReader(request.Body)
    if err == nil {
        _, err = handler.controller.Handle(tlv8)
    }
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        response.WriteHeader(http.StatusNoContent)
    }
}