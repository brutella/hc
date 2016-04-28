package pair

import (
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/util"
	"testing"
)

// Tests the pairing key verification
func TestInvalidPublicKey(t *testing.T) {
	storage, err := util.NewTempFileStorage()

	if err != nil {
		t.Fatal(err)
	}

	database := db.NewDatabaseWithStorage(storage)
	bridge, err := hap.NewSecuredDevice("Macbook Bridge", "001-02-003", database)

	if err != nil {
		t.Fatal(err)
	}

	context := hap.NewContextForSecuredDevice(bridge)

	controller := NewVerifyServerController(database, context)

	client, _ := hap.NewDevice("HomeKit Client", database)
	clientController := NewVerifyClientController(client, database)

	req := clientController.InitialKeyVerifyRequest()
	reqContainer, err := util.NewTLV8ContainerFromReader(req)

	if err != nil {
		t.Fatal(err)
	}

	reqContainer.SetByte(TagPublicKey, byte(0x01))
	// 1) C -> S
	if _, err = HandleReaderForHandler(reqContainer.BytesBuffer(), controller); err != errInvalidClientKeyLength {
		t.Fatal("expected invalid client key length error")
	}
}

// Tests the pairing key verification
func TestPairVerifyIntegration(t *testing.T) {
	storage, err := util.NewTempFileStorage()

	if err != nil {
		t.Fatal(err)
	}

	database := db.NewDatabaseWithStorage(storage)
	bridge, err := hap.NewSecuredDevice("Macbook Bridge", "001-02-003", database)

	if err != nil {
		t.Fatal(err)
	}

	context := hap.NewContextForSecuredDevice(bridge)
	controller := NewVerifyServerController(database, context)

	clientDatabase, _ := db.NewTempDatabase()
	bridgeEntity := db.NewEntity(bridge.Name(), bridge.PublicKey(), nil)
	err = clientDatabase.SaveEntity(bridgeEntity)

	if err != nil {
		t.Fatal(err)
	}

	client, _ := hap.NewDevice("HomeKit Client", clientDatabase)
	clientEntity := db.NewEntity(client.Name(), client.PublicKey(), nil)
	err = database.SaveEntity(clientEntity)

	if err != nil {
		t.Fatal(err)
	}

	clientController := NewVerifyClientController(client, clientDatabase)

	tlvVerifyStepStartRequest := clientController.InitialKeyVerifyRequest()
	// 1) C -> S
	tlvVerifyStepStartResponse, err := HandleReaderForHandler(tlvVerifyStepStartRequest, controller)

	if err != nil {
		t.Fatal(err)
	}

	// 2) S -> C
	tlvFinishRequest, err := HandleReaderForHandler(tlvVerifyStepStartResponse, clientController)

	if err != nil {
		t.Fatal(err)
	}

	// 3) C -> S
	tlvFinishRespond, err := HandleReaderForHandler(tlvFinishRequest, controller)

	if err != nil {
		t.Fatal(err)
	}

	// 4) S -> C
	response, err := HandleReaderForHandler(tlvFinishRespond, clientController)

	if err != nil {
		t.Fatal(err)
	}
	if response != nil {
		t.Fatal(response)
	}
}
