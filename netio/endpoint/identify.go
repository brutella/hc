package endpoint

import (
	"github.com/brutella/hc/netio"
	"github.com/brutella/log"
	"net/http"
)

// Identify handles the unencrypted /identify endpoint by calling IdentifyAccessory() on the IdentifyHandler
type Identify struct {
	http.Handler
	handler netio.IdentifyHandler
}

// NewIdentify returns an object which serves the /identify endpoint
func NewIdentify(h netio.IdentifyHandler) *Identify {
	return &Identify{handler: h}
}

func (i *Identify) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Printf("[VERB] %v POST /identify", request.RemoteAddr)
	i.handler.IdentifyAccessory()
	response.WriteHeader(http.StatusNoContent)
}
