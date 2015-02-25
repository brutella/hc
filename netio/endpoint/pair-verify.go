package endpoint

import (
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/netio/pair"
	"github.com/brutella/log"

	"io"
	"net/http"
)

// Handles the /pair-verify endpoint and returns TLV8 encoded data
//
// This endoint is session based and handles requests based on their connections.
// Which means that there is one pair verify controller for every connection.
// This is required to support simultaneous verification connections.
type PairVerify struct {
	http.Handler
	context  netio.HAPContext
	database db.Database
}

func NewPairVerify(context netio.HAPContext, database db.Database) *PairVerify {
	handler := PairVerify{
		context:  context,
		database: database,
	}

	return &handler
}

func (handler *PairVerify) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("[VERB] POST /pair-verify")
	response.Header().Set("Content-Type", netio.HTTPContentTypePairingTLV8)

	key := handler.context.GetConnectionKey(request)
	session := handler.context.Get(key).(netio.Session)
	controller := session.PairVerifyHandler()
	if controller == nil {
		log.Println("[VERB] Create new pair verify controller")
		controller = pair.NewVerifyServerController(handler.database, handler.context)
		session.SetPairVerifyHandler(controller)
	}

	res, err := pair.HandleReaderForHandler(request.Body, controller)

	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		io.Copy(response, res)
		// Setup secure session
		if controller.KeyVerified() == true {
			// Verification is done
			// Switch to secure session
			secureSession, err := crypto.NewSecureSessionFromSharedKey(controller.SharedKey())
			if err != nil {
				log.Println("[ERRO] Could not setup secure session.", err)
			} else {
				log.Println("[VERB] Setup secure session")
			}
			session.SetCryptographer(secureSession)
		}
	}
}
