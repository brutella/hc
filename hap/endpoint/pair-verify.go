package endpoint

import (
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/hap/pair"
	"github.com/brutella/hc/util"
	"github.com/brutella/log"

	"io"
	"net/http"
)

// PairVerify handles the /pair-verify endpoint and returns TLV8 encoded data
//
// This endoint is session based and handles requests based on their connections.
// Which means that there is one pair verify controller for every connection.
// This is required to support simultaneous verification connections.
type PairVerify struct {
	http.Handler
	context  hap.Context
	database db.Database
}

// NewPairVerify returns a new endpoint for pair verify endpoint
func NewPairVerify(context hap.Context, database db.Database) *PairVerify {
	endpoint := PairVerify{
		context:  context,
		database: database,
	}

	return &endpoint
}

func (endpoint *PairVerify) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Printf("[VERB] %v POST /pair-verify", request.RemoteAddr)
	response.Header().Set("Content-Type", hap.HTTPContentTypePairingTLV8)

	key := endpoint.context.GetConnectionKey(request)
	session := endpoint.context.Get(key).(hap.Session)
	ctlr := session.PairVerifyHandler()
	if ctlr == nil {
		log.Println("[VERB] Create new pair verify controller")
		ctlr = pair.NewVerifyServerController(endpoint.database, endpoint.context)
		session.SetPairVerifyHandler(ctlr)
	}

	var err error
	var in util.Container
	var out util.Container
	var secSession crypto.Cryptographer

	if in, err = util.NewTLV8ContainerFromReader(request.Body); err == nil {
		out, err = ctlr.Handle(in)
	}

	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		io.Copy(response, out.BytesBuffer())

		// When key verification is done, switch to a secure session
		// based on the negotiated shared session key
		b := out.GetByte(pair.TagSequence)
		switch pair.VerifyStepType(b) {
		case pair.VerifyStepFinishResponse:
			if secSession, err = crypto.NewSecureSessionFromSharedKey(ctlr.SharedKey()); err == nil {
				log.Println("[VERB] Setup secure session")
				session.SetCryptographer(secSession)
			} else {
				log.Println("[ERRO] Could not setup secure session.", err)
			}
		}
	}
}
