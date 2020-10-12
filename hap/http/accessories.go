package http

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/log"

	"net/http"
)

func (srv *Server) Accessories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case hap.MethodGET:
		log.Debug.Printf("%v GET /accessories", r.RemoteAddr)

		srv.mutex.Lock()
		if err := WriteJSON(w, r, srv.container); err != nil {
			log.Info.Println(err)
		}
		srv.mutex.Unlock()

	default:
		log.Debug.Println("Cannot handle HTTP method", r.Method)
	}
}
