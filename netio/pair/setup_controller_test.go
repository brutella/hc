package pair

import (
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/util"
	"testing"
)

// Tests the pairing setup
func TestPairingIntegration(t *testing.T) {
	storage, err := util.NewTempFileStorage()

	if err != nil {
		t.Fatal(err)
	}

	database := db.NewDatabaseWithStorage(storage)
	bridge, err := netio.NewSecuredDevice("Macbook Bridge", "001-02-003", database)

	if err != nil {
		t.Fatal(err)
	}

	controller, err := NewSetupServerController(bridge, database)

	if err != nil {
		t.Fatal(err)
	}

	clientDatabase, _ := db.NewTempDatabase()
	client, _ := netio.NewDevice("Client", clientDatabase)
	clientController := NewSetupClientController("001-02-003", client, clientDatabase)
	pairStartRequest := clientController.InitialPairingRequest()

	// 1) C -> S
	pairStartResponse, err := HandleReaderForHandler(pairStartRequest, controller)

	if err != nil {
		t.Fatal(err)
	}

	// 2) S -> C
	pairVerifyRequest, err := HandleReaderForHandler(pairStartResponse, clientController)

	if err != nil {
		t.Fatal(err)
	}

	// 3) C -> S
	pairVerifyResponse, err := HandleReaderForHandler(pairVerifyRequest, controller)

	if err != nil {
		t.Fatal(err)
	}

	// 4) S -> C
	pairKeyRequest, err := HandleReaderForHandler(pairVerifyResponse, clientController)

	if err != nil {
		t.Fatal(err)
	}

	// 5) C -> S
	pairKeyRespond, err := HandleReaderForHandler(pairKeyRequest, controller)

	if err != nil {
		t.Fatal(err)
	}

	// 6) S -> C
	request, err := HandleReaderForHandler(pairKeyRespond, clientController)
	if err != nil {
		t.Fatal(err)
	}

	if request != nil {
		t.Fatal(request)
	}
}
