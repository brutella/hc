package endpoint

import (
	"github.com/brutella/hc/netio"
	"github.com/brutella/log"

	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

// Characteristics handles the /characteristics endpoint
//
// This endpoint is not session based and the same for all connections because
// the encryption/decryption is handled by the connection automatically.
type Characteristics struct {
	http.Handler

	controller netio.CharacteristicsHandler
	mutex      *sync.Mutex
	context    netio.HAPContext
}

// NewCharacteristics returns a new handler for characteristics endpoint
func NewCharacteristics(context netio.HAPContext, c netio.CharacteristicsHandler, mutex *sync.Mutex) *Characteristics {
	handler := Characteristics{
		controller: c,
		mutex:      mutex,
		context:    context,
	}

	return &handler
}

func (handler *Characteristics) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	var res io.Reader
	var err error

	handler.mutex.Lock()
	switch request.Method {
	case netio.MethodGET:
		log.Printf("[VERB] %v GET /characteristics", request.RemoteAddr)
		request.ParseForm()
		res, err = handler.controller.HandleGetCharacteristics(request.Form)
	case netio.MethodPUT:
		log.Printf("[VERB] %v PUT /characteristics", request.RemoteAddr)
		session := handler.context.GetSessionForRequest(request)
		conn := session.Connection()
		err = handler.controller.HandleUpdateCharacteristics(request.Body, conn)
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
