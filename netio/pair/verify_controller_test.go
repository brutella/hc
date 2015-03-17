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
	storage, err := common.NewTempFileStorage()
	assert.Nil(t, err)
	database := db.NewDatabaseWithStorage(storage)
	info := netio.NewBridgeInfo("Macbook Bridge", "001-02-003", "Matthias H.", storage)
	bridge, err := netio.NewBridge(info, database)
	assert.Nil(t, err)
	context := netio.NewContextForBridge(bridge)

	controller := NewVerifyServerController(database, context)

	client, _ := netio.NewClient("HomeKit Client", database)
	client_controller := NewVerifyClientController(client, database)

	req := client_controller.InitialKeyVerifyRequest()
	req_tlv, err := common.NewTLV8ContainerFromReader(req)
	assert.Nil(t, err)
	req_tlv.SetByte(TagPublicKey, byte(0x01))
	// 1) C -> S
	_, err = HandleReaderForHandler(req_tlv.BytesBuffer(), controller)
	assert.Equal(t, err, ErrInvalidClientKeyLength)
}

// Tests the pairing key verification
func TestPairVerifyIntegration(t *testing.T) {
	storage, err := common.NewTempFileStorage()
	assert.Nil(t, err)
	database := db.NewDatabaseWithStorage(storage)
	info := netio.NewBridgeInfo("Macbook Bridge", "001-02-003", "Matthias H.", storage)
	bridge, err := netio.NewBridge(info, database)
	assert.Nil(t, err)
	context := netio.NewContextForBridge(bridge)
	controller := NewVerifyServerController(database, context)

	client_database, _ := db.NewTempDatabase()
	bridge_entity := db.NewEntity(bridge.Id(), bridge.PublicKey(), nil)
	err = client_database.SaveEntity(bridge_entity)
	assert.Nil(t, err)

	client, _ := netio.NewClient("HomeKit Client", client_database)
	client_entity := db.NewEntity(client.Id(), client.PublicKey(), nil)
	err = database.SaveEntity(client_entity)
	assert.Nil(t, err)

	client_controller := NewVerifyClientController(client, client_database)

	tlvVerifyStepStartRequest := client_controller.InitialKeyVerifyRequest()
	// 1) C -> S
	tlvVerifyStepStartResponse, err := HandleReaderForHandler(tlvVerifyStepStartRequest, controller)
	assert.Nil(t, err)

	// 2) S -> C
	tlvFinishRequest, err := HandleReaderForHandler(tlvVerifyStepStartResponse, client_controller)
	assert.Nil(t, err)

	// 3) C -> S
	tlvFinishRespond, err := HandleReaderForHandler(tlvFinishRequest, controller)
	assert.Nil(t, err)

	// 4) S -> C
	response, err := HandleReaderForHandler(tlvFinishRespond, client_controller)
	assert.Nil(t, err)
	assert.Nil(t, response)
}
