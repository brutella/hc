package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"

	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests the pairing key verification
func TestInvalidPublicKey(t *testing.T) {
	storage, err := util.NewTempFileStorage()
	assert.Nil(t, err)
	database := db.NewDatabaseWithStorage(storage)
	bridge, err := netio.NewSecuredDevice("Macbook Bridge", "001-02-003", database)
	assert.Nil(t, err)
	context := netio.NewContextForSecuredDevice(bridge)

	controller := NewVerifyServerController(database, context)

	client, _ := netio.NewDevice("HomeKit Client", database)
	clientController := NewVerifyClientController(client, database)

	req := clientController.InitialKeyVerifyRequest()
	reqContainer, err := util.NewTLV8ContainerFromReader(req)
	assert.Nil(t, err)
	reqContainer.SetByte(TagPublicKey, byte(0x01))
	// 1) C -> S
	_, err = HandleReaderForHandler(reqContainer.BytesBuffer(), controller)
	assert.Equal(t, err, errInvalidClientKeyLength)
}

// Tests the pairing key verification
func TestPairVerifyIntegration(t *testing.T) {
	storage, err := util.NewTempFileStorage()
	assert.Nil(t, err)
	database := db.NewDatabaseWithStorage(storage)
	bridge, err := netio.NewSecuredDevice("Macbook Bridge", "001-02-003", database)
	assert.Nil(t, err)
	context := netio.NewContextForSecuredDevice(bridge)
	controller := NewVerifyServerController(database, context)

	clientDatabase, _ := db.NewTempDatabase()
	bridgeEntity := db.NewEntity(bridge.Name(), bridge.PublicKey(), nil)
	err = clientDatabase.SaveEntity(bridgeEntity)
	assert.Nil(t, err)

	client, _ := netio.NewDevice("HomeKit Client", clientDatabase)
	clientEntity := db.NewEntity(client.Name(), client.PublicKey(), nil)
	err = database.SaveEntity(clientEntity)
	assert.Nil(t, err)

	clientController := NewVerifyClientController(client, clientDatabase)

	tlvVerifyStepStartRequest := clientController.InitialKeyVerifyRequest()
	// 1) C -> S
	tlvVerifyStepStartResponse, err := HandleReaderForHandler(tlvVerifyStepStartRequest, controller)
	assert.Nil(t, err)

	// 2) S -> C
	tlvFinishRequest, err := HandleReaderForHandler(tlvVerifyStepStartResponse, clientController)
	assert.Nil(t, err)

	// 3) C -> S
	tlvFinishRespond, err := HandleReaderForHandler(tlvFinishRequest, controller)
	assert.Nil(t, err)

	// 4) S -> C
	response, err := HandleReaderForHandler(tlvFinishRespond, clientController)
	assert.Nil(t, err)
	assert.Nil(t, response)
}
