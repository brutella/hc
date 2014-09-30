package server

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap"
    "io"
    "io/ioutil"
)

type CharacteristicsHandler struct {
    http.Handler
    
    controller *CharacteristicController
    context *hap.Context
}

func NewCharacteristicsHandler(c *CharacteristicController, context *hap.Context) *CharacteristicsHandler {
    handler := CharacteristicsHandler{
                controller: c,
                context: context,
            }
    
    return &handler
}

func (handler *CharacteristicsHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    
    var res io.Reader
    var err error
    switch request.Method {
    case MethodGET:
        response.Header().Set("Content-Type", hap.HTTPContentTypeHAPJson)
        fmt.Println("GET /characteristics")
        request.ParseForm()
        res, err = handler.controller.HandleGetCharacteristics(request.Form)
    case MethodPUT:
        fmt.Println("PUT /characteristics")
        // no response
        err = handler.controller.HandlePutCharacteristics(request.Body)
    default:
        fmt.Println("Cannot handle HTTP method", request.Method)
    }
    
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        if res != nil {
            bytes, _ := ioutil.ReadAll(res)
            fmt.Println("<-  JSON:", string(bytes))
            response.Write(bytes)
        }
    }
}