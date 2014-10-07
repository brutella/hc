package server

import(
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/netio/pair"
    "github.com/brutella/hap/netio/endpoint"
    "github.com/brutella/hap/netio/controller"
    
    "net/http"
    "fmt"
    "strconv"
    "os"
    "os/signal"
)

type Server interface {
}

type hkServer struct {
    model *model.Model
    context netio.HAPContext
    database db.Database
    bridge *netio.Bridge
    mux *http.ServeMux
    port int
}

func NewServer(c netio.HAPContext, d db.Database, m *model.Model, b *netio.Bridge, port int) *hkServer {
    s := hkServer{
        context: c, 
        database: d, 
        model: m, 
        bridge: b,
        mux: http.NewServeMux(),
        port: port,
    }
    
    s.setupEndpoints()
    
    return &s
}

func (s *hkServer) ListenAndServe() error {
    s.teardownOnExit()
    return netio.ListenAndServe(s.addrString(), s.mux, s.context)
}

func (s *hkServer) Teardown() {
    for _, c := range s.context.ActiveConnection() {
        c.Close()
    }
}

func (s *hkServer) DNSSDCommand() string {
    hostname, _ := os.Hostname()
    return fmt.Sprintf("dns-sd -P %s _hap local %s %s 192.168.0.14 pv=1.0 id=%s c#=1 s#=1 sf=1 ff=0 md=%s\n", s.bridge.Name(), s.portString(), hostname, s.bridge.Id(), s.bridge.Name())
}

func (s *hkServer) teardownOnExit() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, os.Kill)
    
    go func() {
        select {
        case <- c:
            s.Teardown()
            os.Exit(1)
        }
	}()
}

func (s *hkServer) portString() string {
    return strconv.Itoa(s.port)
}

func (s *hkServer) addrString() string {
    return ":" + s.portString()
}

func (s *hkServer) setupEndpoints() {
    model_controller           := controller.NewModelController(s.model)
    characteristics_controller := controller.NewCharacteristicController(s.model)
    pairing_controller         := pair.NewPairingController(s.database)
    
    s.mux.Handle("/pair-setup", endpoint.NewPairSetup(s.bridge, s.database, s.context))
    s.mux.Handle("/pair-verify", endpoint.NewPairVerify(s.context, s.database))
    s.mux.Handle("/accessories", endpoint.NewAccessories(model_controller))
    s.mux.Handle("/characteristics", endpoint.NewCharacteristics(characteristics_controller))
    s.mux.Handle("/pairings", endpoint.NewPairing(pairing_controller))
}