package pair

import (
    "github.com/brutella/hap"
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

// Tests the pairing key verification
func TestPairVerifyIntegration(t *testing.T) {
    bridge, err := hap.NewBridge("HAP Test", "123-45-678")
    assert.Nil(t, err)
    
    storage, err := hap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := hap.NewContext(storage)
    controller, err := NewVerifyServerController(context, bridge)
    assert.Nil(t, err)
    
    name := "UnitTest"
    client_controller := NewVerifyClientController(context, bridge, name)
    context.SaveClient(hap.NewClient(name,client_controller.LTPK)) // make LTPK available to server
    
    tlvVerifyStartRequest := client_controller.InitialKeyVerifyRequest()
    // 1) C -> S
    tlvVerifyStartRespond, err := controller.Handle(tlvVerifyStartRequest)
    assert.Nil(t, err)
    // 2) S -> C
    tlvFinishRequest, err := client_controller.Handle(tlvVerifyStartRespond)
    assert.Nil(t, err)
    // 3) C -> S
    tlvFinishRespond, err := controller.Handle(tlvFinishRequest)
    assert.Nil(t, err)
    assert.True(t, len(context.SharedKey) > 0)
    
    // 4) S -> C 
    response, err := client_controller.Handle(tlvFinishRespond)
    assert.Nil(t, err)
    assert.Nil(t, response)
} 