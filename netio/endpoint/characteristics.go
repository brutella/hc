package endpoint

import (
	"github.com/brutella/hap/netio"
	"github.com/brutella/hap/netio/controller"
	"github.com/brutella/log"

	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

// Handles the /characteristics endpoint
//
// This endpoint is not session based and the same for all connections because
// the encryption/decryption is handled by the connection automatically.
type Characteristics struct {
	http.Handler

	controller *controller.CharacteristicController
	mutex      *sync.Mutex
}

func NewCharacteristics(c *controller.CharacteristicController, mutex *sync.Mutex) *Characteristics {
	handler := Characteristics{
		controller: c,
		mutex:      mutex,
	}

	return &handler
}

func (handler *Characteristics) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	var res io.Reader
	var err error

	handler.mutex.Lock()
	switch request.Method {
	case netio.MethodGET:
		log.Println("[VERB] GET /characteristics")
		request.ParseForm()
		res, err = handler.controller.HandleGetCharacteristics(request.Form)
	case netio.MethodPUT:
		log.Println("[VERB] PUT /characteristics")
		err = handler.controller.HandleUpdateCharacteristics(request.Body)
	default:
		log.Println("[WARN] Cannot handle HTTP method", request.Method)
	}
	handler.mutex.Unlock()

	if err != nil {
		log.Println("[ERRO]", err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		if res != nil {
			response.Header().Set("Content-Type", netio.HTTPContentTypeHAPJson)
			wr := netio.NewChunkedWriter(response, 2048)
			b, _ := ioutil.ReadAll(res)
			wr.Write(b)
		} else {
			response.WriteHeader(http.StatusNoContent)
		}
	}
}
