package endpoint

import (
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/event"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/netio/pair"
	"github.com/brutella/hc/util"
	"github.com/brutella/log"

	"io"
	"net/http"
)

// PairSetup handles the /pair-setup endpoint and returns TLV8 encoded data.
//
// This endoint is session based and handles requests based on their connections.
// Which means that there is one pair setup controller for every connection.
// This is required to support simultaneous pairing connections.
//
// When pairing finished, the DevicePaired event is sent using an event emitter.
type PairSetup struct {
	http.Handler

	device   netio.SecuredDevice
	database db.Database
	context  netio.HAPContext
	emitter  event.Emitter
}

// NewPairSetup returns a new handler for pairing endpoint
func NewPairSetup(context netio.HAPContext, device netio.SecuredDevice, database db.Database, emitter event.Emitter) *PairSetup {
	endpoint := PairSetup{
		device:   device,
		database: database,
		context:  context,
		emitter:  emitter,
	}

	return &endpoint
}

func (endpoint *PairSetup) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Printf("[VERB] %v POST /pair-setup", request.RemoteAddr)
	response.Header().Set("Content-Type", netio.HTTPContentTypePairingTLV8)

	var err error
	var in util.Container
	var out util.Container

	key := endpoint.context.GetConnectionKey(request)
	session := endpoint.context.Get(key).(netio.Session)
	ctrl := session.PairSetupHandler()
	if ctrl == nil {
		log.Println("[VERB] Create new pair setup controller")

		if ctrl, err = pair.NewSetupServerController(endpoint.device, endpoint.database); err != nil {
			log.Println(err)
		}

		session.SetPairSetupHandler(ctrl)
	}

	if in, err = util.NewTLV8ContainerFromReader(request.Body); err == nil {
		out, err = ctrl.Handle(in)
	}

	if err != nil {
		log.Println("[ERRO]", err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		io.Copy(response, out.BytesBuffer())

		// Send event when key exchange is done
		b := out.GetByte(pair.TagSequence)
		switch pair.PairStepType(b) {
		case pair.PairStepKeyExchangeResponse:
			endpoint.emitter.Emit(event.DevicePaired{})
		}
	}
}
