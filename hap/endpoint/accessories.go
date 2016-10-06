package endpoint

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/log"

	"io/ioutil"
	"net/http"
	"sync"
)

// Accessories handles the /accessories endpoint and returns all accessories as JSON
//
// This endpoint is not session based and the same for all connections because
// the encryption/decryption is handled by the connection automatically.
type Accessories struct {
	http.Handler

	controller hap.AccessoriesHandler
	mutex      *sync.Mutex
}

// NewAccessories returns a new handler for accessories endpoint
func NewAccessories(c hap.AccessoriesHandler, mutex *sync.Mutex) *Accessories {
	handler := Accessories{
		controller: c,
		mutex:      mutex,
	}

	return &handler
}

func (handler *Accessories) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Debug.Printf("%v GET /accessories", request.RemoteAddr)
	response.Header().Set("Content-Type", hap.HTTPContentTypeHAPJson)

	handler.mutex.Lock()
	res, err := handler.controller.HandleGetAccessories(request.Body)
	handler.mutex.Unlock()

	if err != nil {
		log.Info.Panic(err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		// Write the data in chunks of 2048 bytes
		// http.ResponseWriter should do this already, but crashes because of an unkown reason
		wr := hap.NewChunkedWriter(response, 2048)
		b, _ := ioutil.ReadAll(res)
		log.Debug.Println(string(b))
		_, err := wr.Write(b)
		if err != nil {
			log.Info.Panic(err)
		}
	}
}
