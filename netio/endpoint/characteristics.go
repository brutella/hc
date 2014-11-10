package endpoint

import(
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/netio/controller"
    "github.com/brutella/log"
    
    "net/http"
    "io"
    "io/ioutil"
)

// Handles the /characteristics endpoint
//
// This endpoint is not session based and the same for all connections because
// the encryption/decryption is handled by the connection automatically.
type Characteristics struct {
    http.Handler
    
    controller *controller.CharacteristicController
}

func NewCharacteristics(c *controller.CharacteristicController) *Characteristics {
    handler := Characteristics{
                controller: c,
            }
    
    return &handler
}

func (handler *Characteristics) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    var res io.Reader
    var err error
    switch request.Method {
    case netio.MethodGET:
        log.Println("[INFO] GET /characteristics")
        request.ParseForm()
        res, err = handler.controller.HandleGetCharacteristics(request.Form)
    case netio.MethodPUT:
        log.Println("[INFO] PUT /characteristics")
        err = handler.controller.HandleUpdateCharacteristics(request.Body)
    default:
        log.Println("[WARN] Cannot handle HTTP method", request.Method)
    }
    
    
    if err != nil {
        log.Println("[ERROR]", err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        if res != nil {
            bytes, _ := ioutil.ReadAll(res)
            response.Header().Set("Content-Type", netio.HTTPContentTypeHAPJson)
            log.Println("[INFO] <-  JSON:", string(bytes))
            response.Write(bytes)
        } else {
            response.WriteHeader(http.StatusNoContent)
        }
    }
}