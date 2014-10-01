package server

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap"
    "io/ioutil"
)

type AccessoriesHandler struct {
    http.Handler
    
    controller *ModelController
    context *hap.Context
}

func NewAccessoriesHandler(c *ModelController, context *hap.Context) *AccessoriesHandler {
    handler := AccessoriesHandler{
                controller: c,
                context: context,
            }
    
    return &handler
}

func (handler *AccessoriesHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("GET /accessories")
    response.Header().Set("Content-Type", hap.HTTPContentTypeHAPJson)
    
    res, err := handler.controller.HandleGetAccessories(request.Body)
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        bytes, _ := ioutil.ReadAll(res)
        fmt.Println("<-  JSON:", string(bytes))
        response.Write(bytes)
    }
}