package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"

	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// Tests the pairing setup
func TestPairingIntegration(t *testing.T) {
	storage, err := common.NewFileStorage(os.TempDir())
	assert.Nil(t, err)
	database := db.NewDatabaseWithStorage(storage)

	info := netio.NewBridgeInfo("Macbook Bridge", "001-02-003", "Matthias H.", storage)
	bridge, err := netio.NewBridge(info)
	assert.Nil(t, err)

	controller, err := NewSetupServerController(bridge, database)
	assert.Nil(t, err)

	client_controller := NewSetupClientController(bridge, "HomeKit Client")
	pairStartRequest := client_controller.InitialPairingRequest()

	// 1) C -> S
	pairStartResponse, err := HandleReaderForHandler(pairStartRequest, controller)
	assert.Nil(t, err)

	// 2) S -> C
	pairVerifyRequest, err := HandleReaderForHandler(pairStartResponse, client_controller)
	assert.Nil(t, err)

	// 3) C -> S
	pairVerifyResponse, err := HandleReaderForHandler(pairVerifyRequest, controller)
	assert.Nil(t, err)

	// 4) S -> C
	pairKeyRequest, err := HandleReaderForHandler(pairVerifyResponse, client_controller)
	assert.Nil(t, err)

	// 5) C -> S
	pairKeyRespond, err := HandleReaderForHandler(pairKeyRequest, controller)
	assert.Nil(t, err)

	// 6) S -> C
	request, err := HandleReaderForHandler(pairKeyRespond, client_controller)
	assert.Nil(t, err)
	assert.Nil(t, request)
}
