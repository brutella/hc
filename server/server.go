package server

import(
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/netio/pair"
    "github.com/brutella/hap/netio/endpoint"
    "github.com/brutella/hap/netio/controller"
    
    "net"
    "net/http"
    "fmt"
    "os"
    "os/signal"
)

func RandomPort() int {
    return 1234
}

type Server interface {
    ListenAndServe() error
}

type ServerExitFunc func()
type hkServer struct {
    model *model.Model
    context netio.HAPContext
    database db.Database
    bridge *netio.Bridge
    mux *http.ServeMux
    port string
    exitFunc ServerExitFunc
}

func NewServer(c netio.HAPContext, d db.Database, m *model.Model, b *netio.Bridge) *hkServer {
    s := hkServer{
        context: c, 
        database: d, 
        model: m, 
        bridge: b,
        mux: http.NewServeMux(),
    }
    
    s.setupEndpoints()
    
    return &s
}

func (s *hkServer) OnExit(fn ServerExitFunc) {
    s.exitFunc = fn
}

func (s *hkServer) ListenAndServe() error {
    s.teardownOnExit()
    
    
    return s.listenAndServe(s.addrString(), s.mux, s.context)
}

func (s *hkServer) Exit() {
    for _, c := range s.context.ActiveConnection() {
        c.Close()
    }
    
    if s.exitFunc != nil {
        s.exitFunc()
    }
}

func (s *hkServer) dnssdCommand() string {
    hostname, _ := os.Hostname()
    return fmt.Sprintf("dns-sd -P %s _hap local %s %s 192.168.0.14 pv=1.0 id=%s c#=1 s#=1 sf=1 ff=0 md=%s\n", s.bridge.Name(), s.port,  hostname, s.bridge.Id(), s.bridge.Name())
}

func (s *hkServer) listenAndServe(addr string, handler http.Handler, context netio.HAPContext) error {
    server := http.Server{Addr: addr, Handler:handler}
    // os gives us a free port when port is "" 
    ln, err := net.Listen("tcp", "")
    if err != nil {
        return err
    }
    
    listener := netio.NewTCPHAPListener(ln.(*net.TCPListener), context)
    
    s.port = ExtractPort(ln.Addr())
    
    fmt.Println(s.dnssdCommand())
    
    return server.Serve(listener)
}

func (s *hkServer) teardownOnExit() {
    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, os.Kill)
    
    go func() {
        select {
        case <- c:
            s.Exit()
            os.Exit(1)
        }
	}()
}

func (s *hkServer) addrString() string {
    return ":" + s.port
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