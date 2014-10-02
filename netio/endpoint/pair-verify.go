package endpoint

import(
    "github.com/brutella/hap/netio/pair"
    "github.com/brutella/hap/netio"
    "github.com/brutella/hap/crypto"
    "github.com/brutella/hap/db"
    
    "net/http"
    "fmt"
    "io/ioutil"
)

type PairVerify struct {
    http.Handler
    context netio.HAPContext
    database db.Database
}

func NewPairVerify(context netio.HAPContext, database db.Database) *PairVerify {
    handler := PairVerify{
                context: context,
                database: database,
            }
    
    return &handler
}

func (handler *PairVerify) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    fmt.Println("POST /pair-verify")
    response.Header().Set("Content-Type", netio.HTTPContentTypePairingTLV8)
    
    key := handler.context.GetConnectionKey(request)
    session := handler.context.Get(key).(netio.Session)
    controller := session.PairVerifyHandler()
    if controller == nil {
        fmt.Println("Create new pair verify controller")
        controller = pair.NewVerifyServerController(handler.database, handler.context)
        session.SetPairVerifyHandler(controller)
    }
    
    res, err := pair.HandleReaderForHandler(request.Body, controller)
    
    if err != nil {
        fmt.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
    } else {
        bytes, _ := ioutil.ReadAll(res)
        response.Write(bytes)
        
        // Setup secure session
        if controller.KeyVerified() == true {
            // Verification is done
            // Switch to secure session
            secureSession, err := crypto.NewSecureSessionFromSharedKey(controller.SharedKey())
            if err != nil {
                fmt.Println("Could not setup secure session.", err)
            } else {
                fmt.Println("Setup secure session")
            }
            session.SetCryptographer(secureSession)
        }
    }
}