package endpoint

import (
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/event"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/hap/pair"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/util"

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

	device   hap.SecuredDevice
	database db.Database
	context  hap.Context
	emitter  event.Emitter
}

// NewPairSetup returns a new handler for pairing endpoint
func NewPairSetup(context hap.Context, device hap.SecuredDevice, database db.Database, emitter event.Emitter) *PairSetup {
	endpoint := PairSetup{
		device:   device,
		database: database,
		context:  context,
		emitter:  emitter,
	}

	return &endpoint
}

func (endpoint *PairSetup) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Debug.Printf("%v POST /pair-setup", request.RemoteAddr)
	response.Header().Set("Content-Type", hap.HTTPContentTypePairingTLV8)

	var err error
	var in util.Container
	var out util.Container

	key := endpoint.context.GetConnectionKey(request)
	session := endpoint.context.Get(key).(hap.Session)
	ctrl := session.PairSetupHandler()
	if ctrl == nil {
		log.Debug.Println("Create new pair setup controller")

		if ctrl, err = pair.NewSetupServerController(endpoint.device, endpoint.database); err != nil {
			log.Info.Panic(err)
		}

		session.SetPairSetupHandler(ctrl)
	}

	if in, err = util.NewTLV8ContainerFromReader(request.Body); err == nil {
		out, err = ctrl.Handle(in)
	}

	if err != nil {
		log.Info.Println(err)
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
