package http

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/event"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/hap/controller"
	"github.com/brutella/hc/hap/endpoint"
	"github.com/brutella/hc/hap/pair"
	"github.com/brutella/hc/log"

	"net"
	"net/http"
	"sync"
)

// Server provides a similar interfaces as http.Server to start and stop a TCP server.
type Server interface {
	// ListenAndServe start the server
	ListenAndServe() error

	// Port returns the port on which the server listens to
	Port() string

	// Stop stops the server
	Stop()
}

type Config struct {
	Port      string
	Context   hap.Context
	Database  db.Database
	Container *accessory.Container
	Device    hap.SecuredDevice
	Mutex     *sync.Mutex
	Emitter   event.Emitter
}

type server struct {
	context  hap.Context
	database db.Database
	device   hap.SecuredDevice
	mux      *http.ServeMux

	mutex     *sync.Mutex
	container *accessory.Container

	port        string
	listener    *net.TCPListener
	hapListener *hap.TCPListener

	emitter event.Emitter
}

// NewServer returns a server
func NewServer(c Config) Server {

	// os gives us a free Port when Port is ""
	ln, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Info.Panic(err)
	}

	_, port, _ := net.SplitHostPort(ln.Addr().String())

	s := server{
		context:   c.Context,
		database:  c.Database,
		container: c.Container,
		device:    c.Device,
		mux:       http.NewServeMux(),
		mutex:     c.Mutex,
		listener:  ln.(*net.TCPListener),
		port:      port,
		emitter:   c.Emitter,
	}

	s.setupEndpoints()

	return &s
}

func (s *server) ListenAndServe() error {
	return s.listenAndServe(s.addrString(), s.mux, s.context)
}

func (s *server) Stop() {
	for _, c := range s.context.ActiveConnections() {
		c.Close()
	}
	// Stop listener
	s.hapListener.Close()
}

func (s *server) Port() string {
	return s.port
}

// listenAndServe returns a http.Server to listen on a specific address
func (s *server) listenAndServe(addr string, handler http.Handler, context hap.Context) error {
	server := http.Server{Addr: addr, Handler: handler}
	// Use a TCPListener
	listener := hap.NewTCPListener(s.listener, context)
	s.hapListener = listener
	return server.Serve(listener)
}

func (s *server) addrString() string {
	return ":" + s.port
}

// setupEndpoints creates controller objects to handle HAP endpoints
func (s *server) setupEndpoints() {
	containerController := controller.NewContainerController(s.container)
	characteristicsController := controller.NewCharacteristicController(s.container)
	pairingController := pair.NewPairingController(s.database)

	s.mux.Handle("/pair-setup", endpoint.NewPairSetup(s.context, s.device, s.database, s.emitter))
	s.mux.Handle("/pair-verify", endpoint.NewPairVerify(s.context, s.database))
	s.mux.Handle("/accessories", endpoint.NewAccessories(containerController, s.mutex))
	s.mux.Handle("/characteristics", endpoint.NewCharacteristics(s.context, characteristicsController, s.mutex))
	s.mux.Handle("/pairings", endpoint.NewPairing(pairingController, s.emitter))
	s.mux.Handle("/identify", endpoint.NewIdentify(containerController))
}
