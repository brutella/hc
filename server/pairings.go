package server

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap"
    "io"
    "io/ioutil"
)

type PairingHandler struct {
    http.Handler
    controller *PairingController
    context *hap.Context
}

func NewPairingHandler(controller *PairingController, context *hap.Context) *PairingHandler {
    handler := PairingHandler{
                controller: controller,
                context: context,
            }
    
    return &handler
}

func (handler *PairingHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {    
    var res io.Reader
    var err error
    switch request.Method {
    case MethodDEL:
        fmt.Println("DEL /pairings")
        res, err = handler.controller.HandleDeletePairings(request.Body)
    case MethodPOST:
        fmt.Println("POST /pairings")
        res, err = handler.controller.HandlePostPairings(request.Body)
    default:
        fmt.Println("Cannot handle HTTP method", request.Method)
    }
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        bytes, _ := ioutil.ReadAll(res)
        fmt.Println("<-  JSON:", string(bytes))
        response.Write(bytes)
    }
}