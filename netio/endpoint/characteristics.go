package endpoint

import(
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/netio/controller"
    
    "net/http"
    "fmt"
    "io"
    "io/ioutil"
)

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
        fmt.Println("GET /characteristics")
        request.ParseForm()
        res, err = handler.controller.HandleGetCharacteristics(request.Form)
    case netio.MethodPUT:
        fmt.Println("PUT /characteristics")
        err = handler.controller.HandleUpdateCharacteristics(request.Body)
    default:
        fmt.Println("Cannot handle HTTP method", request.Method)
    }
    
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        if res != nil {
            bytes, _ := ioutil.ReadAll(res)
            response.Header().Set("Content-Type", netio.HTTPContentTypeHAPJson)
            fmt.Println("<-  JSON:", string(bytes))
            response.Write(bytes)
        } else {
            response.WriteHeader(http.StatusNoContent)
        }
    }
}