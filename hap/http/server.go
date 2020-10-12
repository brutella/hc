package http

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/event"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/hap/endpoint"
	"github.com/brutella/hc/hap/pair"
	"github.com/brutella/hc/log"

	"context"
	"net"
	"net/http"
	"sync"
)

type Config struct {
	Port      string
	Context   hap.Context
	Database  db.Database
	Container *accessory.Container
	Device    hap.SecuredDevice
	Mutex     *sync.Mutex
	Emitter   event.Emitter
}

type Server struct {
	context  hap.Context
	database db.Database
	device   hap.SecuredDevice
	Mux      *http.ServeMux

	mutex     *sync.Mutex
	container *accessory.Container

	port     string
	listener *net.TCPListener

	emitter event.Emitter
}

// NewServer returns a server
func NewServer(c Config) *Server {

	// os gives us a free Port when Port is ""
	ln, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Info.Panic(err)
	}

	_, port, _ := net.SplitHostPort(ln.Addr().String())

	s := Server{
		context:   c.Context,
		database:  c.Database,
		container: c.Container,
		device:    c.Device,
		Mux:       http.NewServeMux(),
		mutex:     c.Mutex,
		listener:  ln.(*net.TCPListener),
		port:      port,
		emitter:   c.Emitter,
	}

	s.setupEndpoints()

	return &s
}

func testable(c Config) *Server {
	s := Server{
		context:   c.Context,
		database:  c.Database,
		container: c.Container,
		device:    c.Device,
		Mux:       http.NewServeMux(),
		mutex:     c.Mutex,
		emitter:   c.Emitter,
	}

	s.setupEndpoints()

	return &s
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		for _, c := range s.context.ActiveConnections() {
			c.Close()
		}
		// Stop listener
		s.listener.Close()
	}()
	return s.listenAndServe(s.addrString(), s.Mux, s.context)
}

func (s *Server) Port() string {
	return s.port
}

func (s *Server) listenAndServe(addr string, handler http.Handler, context hap.Context) error {
	server := http.Server{Addr: addr, Handler: handler}
	return server.Serve(s)
}

func (s *Server) addrString() string {
	return ":" + s.port
}

// setupEndpoints creates controller objects to handle HAP endpoints
func (s *Server) setupEndpoints() {
	pairingController := pair.NewPairingController(s.database)

	s.Mux.Handle("/pair-setup", endpoint.NewPairSetup(s.context, s.device, s.database, s.emitter))
	s.Mux.Handle("/pair-verify", endpoint.NewPairVerify(s.context, s.database))
	s.Mux.Handle("/accessories", s.Authenticate(http.HandlerFunc(s.Accessories)))
	s.Mux.Handle("/characteristics", s.Authenticate(http.HandlerFunc(s.Characteristics)))
	s.Mux.Handle("/pairings", endpoint.NewPairing(pairingController, s.emitter))
	s.Mux.HandleFunc("/identify", s.Identify)
}

func (srv *Server) getCharacteristic(aid uint64, iid uint64) *characteristic.Characteristic {
	for _, a := range srv.container.Accessories {
		if a.ID == aid {
			for _, s := range a.GetServices() {
				for _, c := range s.GetCharacteristics() {
					if c.ID == iid {
						return c
					}
				}
			}
		}
	}
	return nil
}
