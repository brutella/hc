package handler

import(   
    _"github.com/brutella/hap"
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/netio/controller"
    
    "net/http"
    "io/ioutil"
    "fmt"
)

type AccessoriesHandler struct {
    http.Handler
    
    controller *controller.ModelController
    context netio.Context
}

func NewAccessoriesHandler(c *controller.ModelController, context netio.Context) *AccessoriesHandler {
    handler := AccessoriesHandler{
                controller: c,
                context: context,
            }
    
    return &handler
}

func (handler *AccessoriesHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("GET /accessories")
    response.Header().Set("Content-Type", netio.HTTPContentTypeHAPJson)
    
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