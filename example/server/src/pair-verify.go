package hapserver

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
}

func NewPairVerifyHandler(c *pair.VerifyServerController) *PairVerifyHandler {
    handler := PairVerifyHandler{
                controller: c,
            }
    
    return &handler
}

func (handler *PairVerifyHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("Pair-Verify request")
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