package endpoint

import(   
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/netio/controller"
    "github.com/brutella/log"
    
    "net/http"
    "io/ioutil"
    "sync"
)

// Handles the /accessories endpoint and returns all accessories as JSON
//
// This endpoint is not session based and the same for all connections because
// the encryption/decryption is handled by the connection automatically.
type Accessories struct {
    http.Handler
    
    controller *controller.ContainerController
    mutex *sync.Mutex
}

func NewAccessories(c *controller.ContainerController, mutex *sync.Mutex) *Accessories {
    handler := Accessories{
                controller: c,
                mutex: mutex,
            }
    
    return &handler
}

func (handler *Accessories) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    log.Println("[VERB] GET /accessories")
    response.Header().Set("Content-Type", netio.HTTPContentTypeHAPJson)
    
    handler.mutex.Lock()
    res, err := handler.controller.HandleGetAccessories(request.Body)
    handler.mutex.Unlock()
    
    if err != nil {
        log.Println("[ERRO]", err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        // Write the data in chunks of 2048 bytes
        // http.ResponseWriter should do this already, but crashes because of an unkown reason
        wr := netio.NewChunkedWriter(response, 2048)
        b, _ := ioutil.ReadAll(res)
        _, err := wr.Write(b)
        if err != nil {
            log.Println("[ERRO]", err)
        }
    }
}