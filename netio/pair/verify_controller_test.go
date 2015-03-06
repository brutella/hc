package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"

	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// Tests the pairing key verification
func TestPairVerifyIntegration(t *testing.T) {
	storage, err := common.NewFileStorage(os.TempDir())
	assert.Nil(t, err)
	database := db.NewDatabaseWithStorage(storage)
	info := netio.NewBridgeInfo("Macbook Bridge", "001-02-003", "Matthias H.", storage)
	bridge, err := netio.NewBridge(info)
	assert.Nil(t, err)
	context := netio.NewContextForBridge(bridge)

	controller := NewVerifyServerController(database, context)

	name := "UnitTest"
	client_controller := NewVerifyClientController(bridge, name)
	database.SaveClient(db.NewClient(name, client_controller.LTPK)) // make LTPK available to server

	tlvSequenceVerifyStartRequest := client_controller.InitialKeyVerifyRequest()
	// 1) C -> S
	tlvSequenceVerifyStartResponse, err := HandleReaderForHandler(tlvSequenceVerifyStartRequest, controller)
	assert.Nil(t, err)

	// 2) S -> C
	tlvFinishRequest, err := HandleReaderForHandler(tlvSequenceVerifyStartResponse, client_controller)
	assert.Nil(t, err)

	// 3) C -> S
	tlvFinishRespond, err := HandleReaderForHandler(tlvFinishRequest, controller)
	assert.Nil(t, err)

	// 4) S -> C
	response, err := HandleReaderForHandler(tlvFinishRespond, client_controller)
	assert.Nil(t, err)
	assert.Nil(t, response)
}
