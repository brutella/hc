package endpoint

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/log"
	"net/http"
)

// Identify handles the unencrypted /identify endpoint by calling IdentifyAccessory() on the IdentifyHandler
type Identify struct {
	http.Handler
	handler hap.IdentifyHandler
}

// NewIdentify returns an object which serves the /identify endpoint
func NewIdentify(h hap.IdentifyHandler) *Identify {
	return &Identify{handler: h}
}

func (i *Identify) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Debug.Printf("%v POST /identify", request.RemoteAddr)
	i.handler.IdentifyAccessory()
	response.WriteHeader(http.StatusNoContent)
}
