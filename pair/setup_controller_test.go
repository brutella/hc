package pair

import (
    "github.com/brutella/hap"

    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

// Tests the pairing setup
func TestPairingIntegration(t *testing.T) {
    storage, err := hap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := hap.NewContext(storage)
    
    info := hap.NewBridgeInfo("Macbook Bridge", "001-02-003", "Matthias H.", storage)
    bridge, err := hap.NewBridge(info)
    assert.Nil(t, err)
    
    controller, err := NewSetupServerController(context, bridge)
    assert.Nil(t, err)
    
    client_controller := NewSetupClientController(context, bridge, "HomeKit Client")
    pairStartRequest := client_controller.InitialPairingRequest()
    
    // 1) C -> S
    pairStartRespond, err := controller.HandleReader(pairStartRequest)
    assert.Nil(t, err)
    
    // 2) S -> C
    pairVerifyRequest, err := client_controller.HandleReader(pairStartRespond)
    assert.Nil(t, err)
    
    // 3) C -> S
    pairVerifyRespond, err := controller.HandleReader(pairVerifyRequest)
    assert.Nil(t, err)
    
    // 4) S -> C
    pairKeyRequest, err := client_controller.HandleReader(pairVerifyRespond)
    assert.Nil(t, err)
    
    // 5) C -> S
    pairKeyRespond, err := controller.HandleReader(pairKeyRequest)
    assert.Nil(t, err)
    
    // 6) S -> C
    request, err := client_controller.HandleReader(pairKeyRespond)
    assert.Nil(t, err)
    assert.Nil(t, request)
}  