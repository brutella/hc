package http

import (
	"github.com/brutella/hc/log"

	"net/http"
)

// TODO respond with 400 if accessory is already paired
//
//     HTTP/1.1 400 Bad Request
//     Content-Type: application/hap+json
//     Content-Length: <length>
//     { "status" : -70401 }
func (srv *Server) Identify(w http.ResponseWriter, r *http.Request) {
	log.Debug.Printf("%v POST /identify", r.RemoteAddr)
	for _, a := range srv.container.Accessories {
		a.Identify()
	}
	w.WriteHeader(http.StatusNoContent)
}
