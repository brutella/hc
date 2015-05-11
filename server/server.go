package server

import (
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/model/container"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/netio/controller"
	"github.com/brutella/hc/netio/endpoint"
	"github.com/brutella/hc/netio/pair"

	"log"
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

type hkServer struct {
	context  netio.HAPContext
	database db.Database
	device   netio.SecuredDevice
	mux      *http.ServeMux

	mutex     *sync.Mutex
	container *container.Container

	port        string
	listener    *net.TCPListener
	hapListener *netio.HAPTCPListener
}

// NewServer returns a server
func NewServer(ctx netio.HAPContext, d db.Database, c *container.Container, device netio.SecuredDevice, mutex *sync.Mutex) Server {
	// os gives us a free Port when Port is ""
	ln, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatal(err)
	}

	_, port, _ := net.SplitHostPort(ln.Addr().String())

	s := hkServer{
		context:   ctx,
		database:  d,
		container: c,
		device:    device,
		mux:       http.NewServeMux(),
		mutex:     mutex,
		listener:  ln.(*net.TCPListener),
		port:      port,
	}

	s.setupEndpoints()

	return &s
}

func (s *hkServer) ListenAndServe() error {
	return s.listenAndServe(s.addrString(), s.mux, s.context)
}

func (s *hkServer) Stop() {
	for _, c := range s.context.ActiveConnections() {
		c.Close()
	}
	// Stop listener
	s.hapListener.Close()
}

func (s *hkServer) Port() string {
	return s.port
}

// listenAndServe returns a http.Server to listen on a specific address
func (s *hkServer) listenAndServe(addr string, handler http.Handler, context netio.HAPContext) error {
	server := http.Server{Addr: addr, Handler: handler}
	// Use a HAPTCPListener
	listener := netio.NewHAPTCPListener(s.listener, context)
	s.hapListener = listener
	return server.Serve(listener)
}

func (s *hkServer) addrString() string {
	return ":" + s.port
}

// setupEndpoints creates controller objects to handle HAP endpoints
func (s *hkServer) setupEndpoints() {
	containerController := controller.NewContainerController(s.container)
	characteristicsController := controller.NewCharacteristicController(s.container)
	pairingController := pair.NewPairingController(s.database)

	s.mux.Handle("/pair-setup", endpoint.NewPairSetup(s.context, s.device, s.database))
	s.mux.Handle("/pair-verify", endpoint.NewPairVerify(s.context, s.database))
	s.mux.Handle("/accessories", endpoint.NewAccessories(containerController, s.mutex))
	s.mux.Handle("/characteristics", endpoint.NewCharacteristics(s.context, characteristicsController, s.mutex))
	s.mux.Handle("/pairings", endpoint.NewPairing(pairingController))
	s.mux.Handle("/identify", endpoint.NewIdentify(containerController))
}
