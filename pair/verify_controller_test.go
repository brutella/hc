package pair

import (
    "github.com/brutella/hap"
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

// Tests the pairing key verification
func TestPairVerifyIntegration(t *testing.T) {
    storage, err := hap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := hap.NewContext(storage)
    info := hap.NewBridgeInfo("Macbook Bridge", "001-02-003", "Matthias H.", storage)
    bridge, err := hap.NewBridge(info)
    assert.Nil(t, err)
    
    controller, err := NewVerifyServerController(context, bridge)
    assert.Nil(t, err)
    
    name := "UnitTest"
    client_controller := NewVerifyClientController(context, bridge, name)
    context.SaveClient(hap.NewClient(name,client_controller.LTPK)) // make LTPK available to server
    
    tlvVerifyStartRequest := client_controller.InitialKeyVerifyRequest()
    // 1) C -> S
    tlvVerifyStartRespond, err := HandleReaderForHandler(tlvVerifyStartRequest, controller)
    assert.Nil(t, err)
    
    // 2) S -> C
    tlvFinishRequest, err := HandleReaderForHandler(tlvVerifyStartRespond, client_controller)
    assert.Nil(t, err)
    
    // 3) C -> S
    tlvFinishRespond, err := HandleReaderForHandler(tlvFinishRequest, controller)
    assert.Nil(t, err)
    assert.NotNil(t, context.SecSession) // secure session is established
    
    // 4) S -> C 
    response, err := HandleReaderForHandler(tlvFinishRespond, client_controller)
    assert.Nil(t, err)
    assert.Nil(t, response)
} 