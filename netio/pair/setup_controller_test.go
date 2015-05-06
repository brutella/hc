package pair

import (
	"github.com/brutella/hc/util"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"

	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests the pairing setup
func TestPairingIntegration(t *testing.T) {
	storage, err := util.NewTempFileStorage()
	assert.Nil(t, err)
	database := db.NewDatabaseWithStorage(storage)
	bridge, err := netio.NewSecuredDevice("Macbook Bridge", "001-02-003", database)
	assert.Nil(t, err)

	controller, err := NewSetupServerController(bridge, database)
	assert.Nil(t, err)
	clientDatabase, _ := db.NewTempDatabase()
	client, _ := netio.NewDevice("Client", clientDatabase)
	clientController := NewSetupClientController("001-02-003", client, clientDatabase)
	pairStartRequest := clientController.InitialPairingRequest()

	// 1) C -> S
	pairStartResponse, err := HandleReaderForHandler(pairStartRequest, controller)
	assert.Nil(t, err)

	// 2) S -> C
	pairVerifyRequest, err := HandleReaderForHandler(pairStartResponse, clientController)
	assert.Nil(t, err)

	// 3) C -> S
	pairVerifyResponse, err := HandleReaderForHandler(pairVerifyRequest, controller)
	assert.Nil(t, err)

	// 4) S -> C
	pairKeyRequest, err := HandleReaderForHandler(pairVerifyResponse, clientController)
	assert.Nil(t, err)

	// 5) C -> S
	pairKeyRespond, err := HandleReaderForHandler(pairKeyRequest, controller)
	assert.Nil(t, err)

	// 6) S -> C
	request, err := HandleReaderForHandler(pairKeyRespond, clientController)
	assert.Nil(t, err)
	assert.Nil(t, request)
}
