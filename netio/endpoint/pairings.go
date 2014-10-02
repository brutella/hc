package endpoint

import(
    "net/http"
    "fmt"
    _"github.com/brutella/hap"
    "github.com/brutella/hap/netio/pair"
)

type Pairing struct {
    http.Handler
    
    controller *pair.PairingController
}

func NewPairing(controller *pair.PairingController) *Pairing {
    handler := Pairing{
                controller: controller,
            }
    
    return &handler
}

func (handler *Pairing) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("POST /pairings")
    
    _, err := pair.HandleReaderForHandler(request.Body, handler.controller)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        response.WriteHeader(http.StatusNoContent)
    }
}