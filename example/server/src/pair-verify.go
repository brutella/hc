package hapserver

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap/crypto"
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
    
    if handler.controller.KeyVerificationCompleted() == true {
        // Verification is done
        // Switch to secure session
        secSession, err := crypto.NewSecureSessionFromSharedKey(handler.controller.VerifiedSharedKey())
        if err != nil {
            fmt.Println("Could not setup secure session.", err)
        } else {
            fmt.Println("Setup secure session")
        }
        handler.context.SetSecureSession(secSession) 
    } else {
        handler.context.SetSecureSession(nil)
    }
}