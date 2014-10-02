package endpoint

import(   
    _"github.com/brutella/hap"
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/netio/controller"
    
    "net/http"
    "io/ioutil"
    "fmt"
)

type Accessories struct {
    http.Handler
    
    controller *controller.ModelController
}

func NewAccessories(c *controller.ModelController) *Accessories {
    handler := Accessories{
                controller: c,
            }
    
    return &handler
}

func (handler *Accessories) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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