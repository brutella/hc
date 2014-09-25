package hapserver

import(
    "net/http"
    "fmt"
    "github.com/brutella/hap"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/crypto"
    "io/ioutil"
)

type AccessoriesHandler struct {
    http.Handler
    
    controller *model.ModelController
    context *hap.Context
}

func NewAccessoriesHandler(c *model.ModelController, context *hap.Context) *AccessoriesHandler {
    handler := AccessoriesHandler{
                controller: c,
                context: context,
            }
    
    return &handler
}

func (handler *AccessoriesHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("Get Accessories Request")
    response.Header().Set("Content-Type", hap.HTTPContentTypeHAPJson)
    
    decrypted, err := crypto.Decrypt(request.Body, handler.context)
    
    if err != nil {
        fmt.Println("Decryption failed", err)
        response.WriteHeader(http.StatusInternalServerError)
    }
    
    res, err := handler.controller.HandleGetAccessories(decrypted)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        encrypted, err := crypto.Encrypt(res, handler.context)
        
        if err != nil {
            fmt.Println("Encryption failed", err)
            response.WriteHeader(http.StatusInternalServerError)
        }
        
        bytes, _ := ioutil.ReadAll(encrypted)
        response.Write(bytes)
    }
}