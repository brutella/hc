package endpoint

import (
	"github.com/brutella/log"
	"net/http"
)

// Identify handles the unencrypted /identify endpoint.
type Identify struct {
	http.Handler
}

// NewPairing returns a new handler for pairing enpdoint
func NewIdentify() *Identify {
	return &Identify{}
}

func (handler *Identify) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Printf("[VERB] %v POST /identify", request.RemoteAddr)
	response.WriteHeader(http.StatusNoContent)
}
