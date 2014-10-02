package pair

import (
    "github.com/brutella/hap"
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/netio"

    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

// Tests the pairing setup
func TestPairingIntegration(t *testing.T) {
    storage, err := hap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    database := db.NewManager(storage)
    
    info := netio.NewBridgeInfo("Macbook Bridge", "001-02-003", "Matthias H.", storage)
    bridge, err := netio.NewBridge(info)
    assert.Nil(t, err)
    
    controller, err := NewSetupServerController(bridge, database)
    assert.Nil(t, err)
    
    client_controller := NewSetupClientController(bridge, "HomeKit Client")
    pairStartRequest := client_controller.InitialPairingRequest()
    
    // 1) C -> S
    pairStartRespond, err := HandleReaderForHandler(pairStartRequest, controller)
    assert.Nil(t, err)
    
    // 2) S -> C
    pairVerifyRequest, err := HandleReaderForHandler(pairStartRespond, client_controller)
    assert.Nil(t, err)
    
    // 3) C -> S
    pairVerifyRespond, err := HandleReaderForHandler(pairVerifyRequest, controller)
    assert.Nil(t, err)
    
    // 4) S -> C
    pairKeyRequest, err := HandleReaderForHandler(pairVerifyRespond, client_controller)
    assert.Nil(t, err)
    
    // 5) C -> S
    pairKeyRespond, err := HandleReaderForHandler(pairKeyRequest, controller)
    assert.Nil(t, err)
    
    // 6) S -> C
    request, err := HandleReaderForHandler(pairKeyRespond, client_controller)
    assert.Nil(t, err)
    assert.Nil(t, request)
}  