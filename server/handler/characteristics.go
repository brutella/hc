package handler

import(
    "github.com/brutella/hap"
    "github.com/brutella/hap/server"
    "github.com/brutella/hap/server/controller"
    
    "net/http"
    "fmt"
    "io"
    "io/ioutil"
    "encoding/json"
    "bytes"
)

type CharacteristicsHandler struct {
    http.Handler
    
    controller *controller.CharacteristicController
    context *hap.Context
}

func NewCharacteristicsHandler(c *controller.CharacteristicController, context *hap.Context) *CharacteristicsHandler {
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
    case server.MethodGET:
        fmt.Println("GET /characteristics")
        request.ParseForm()
        aid, cid, err := server.ParseAccessoryAndCharacterId(request.Form.Get("id"))
        chars := handler.controller.HandleGetCharacteristics(aid, cid)
        result, err := json.Marshal(chars)
        if err != nil {
            fmt.Println(err)
        }
        
        var b bytes.Buffer
        b.Write(result)
        res = &b
    case server.MethodPUT:
        fmt.Println("PUT /characteristics")
        
        b, _ := ioutil.ReadAll(request.Body)
        var chars controller.Characteristics
        err := json.Unmarshal(b, &chars)
    
        if err != nil {
            fmt.Println("Could not unmarshal to json", err)
        } else {
            err = handler.controller.HandleUpdateCharacteristics(chars)
        }
    default:
        fmt.Println("Cannot handle HTTP method", request.Method)
    }
    
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        if res != nil {
            bytes, _ := ioutil.ReadAll(res)
            response.Header().Set("Content-Type", server.HTTPContentTypeHAPJson)
            fmt.Println("<-  JSON:", string(bytes))
            response.Write(bytes)
        } else {
            response.WriteHeader(http.StatusNoContent)
        }
    }
}