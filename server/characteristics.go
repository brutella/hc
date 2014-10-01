package server

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap"
    "io"
    "io/ioutil"
    "encoding/json"
    "bytes"
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
        fmt.Println("GET /characteristics")
        request.ParseForm()
        aid, cid, err := ParseAccessoryAndCharacterId(request.Form.Get("id"))
        chars := handler.controller.HandleGetCharacteristics(aid, cid)
        result, err := json.Marshal(chars)
        if err != nil {
            fmt.Println(err)
        }
        
        var b bytes.Buffer
        b.Write(result)
        res = &b
    case MethodPUT:
        fmt.Println("PUT /characteristics")
        
        b, _ := ioutil.ReadAll(request.Body)
        var chars Characteristics
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
            response.Header().Set("Content-Type", hap.HTTPContentTypeHAPJson)
            fmt.Println("<-  JSON:", string(bytes))
            response.Write(bytes)
        } else {
            response.WriteHeader(http.StatusNoContent)
        }
    }
}