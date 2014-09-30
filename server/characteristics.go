package server

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap"
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
    fmt.Println("GET /characteristics")
    fmt.Println("id = ", request.FormValue("id"))
    response.Header().Set("Content-Type", hap.HTTPContentTypeHAPJson)
    
    res, err := handler.controller.HandleGetCharacteristics(request.Form)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        bytes, _ := ioutil.ReadAll(res)
        fmt.Println("<-  JSON:", string(bytes))
        response.Write(bytes)
    }
}