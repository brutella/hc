package endpoint

import(
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/netio/pair"
    "github.com/brutella/hap/netio"
    "github.com/brutella/log"
    
    "io"
    "net/http"
)

// Handles the /pair-setup endpoint and returns TLV8 encoded data
//
// This endoint is session based and handles requests based on their connections.
// Which means that for every unique connection, there will be a new controller
// set up. This is required to support simultaneous pairigin connections.
type PairSetup struct {
    http.Handler
    
    bridge *netio.Bridge
    database db.Database
    context netio.HAPContext
}

func NewPairSetup(bridge *netio.Bridge, database db.Database, context netio.HAPContext) *PairSetup {
    handler := PairSetup{
                bridge: bridge,
                database: database,
                context: context,
            }
    
    return &handler
}

func (handler *PairSetup) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    log.Println("[VERB] POST /pair-setup")
    response.Header().Set("Content-Type", netio.HTTPContentTypePairingTLV8)
    
    key := handler.context.GetConnectionKey(request)
    session := handler.context.Get(key).(netio.Session)
    controller := session.PairSetupHandler()
    if controller == nil {
        log.Println("[VERB] Create new pair setup controller")
        var err error
        controller, err = pair.NewSetupServerController(handler.bridge, handler.database)
        if err != nil {
            log.Println(err)
        }
        
        session.SetPairSetupHandler(controller)
    }
    
    res, err := pair.HandleReaderForHandler(request.Body, controller)
    
    if err != nil {
        log.Println("[ERRO]", err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        io.Copy(response, res)
    }
}