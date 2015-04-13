package endpoint

import (
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/netio/pair"
	"github.com/brutella/log"

	"io"
	"net/http"
)

// PairSetup handles the /pair-setup endpoint and returns TLV8 encoded data.
//
// This endoint is session based and handles requests based on their connections.
// Which means that there is one pair setup controller for every connection.
// This is required to support simultaneous pairing connections.
type PairSetup struct {
	http.Handler

	device   netio.SecuredDevice
	database db.Database
	context  netio.HAPContext
}

// NewPairSetup returns a new handler for pairing endpoint
func NewPairSetup(device netio.SecuredDevice, database db.Database, context netio.HAPContext) *PairSetup {
	handler := PairSetup{
		device:   device,
		database: database,
		context:  context,
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
		controller, err = pair.NewSetupServerController(handler.device, handler.database)
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
