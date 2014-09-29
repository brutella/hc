package hk

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "encoding/json"
)

func TestAccessory(t *testing.T) {
    // serialNumber, modelName, manufacturerName, accessoryName string
    info_service := NewAccessoryInfoService("123-456-789", "Rev1", "Matthias H.", "My Bridge")
    accessory := NewAccessory()
    accessory.AddService(info_service.Service)
    
    result, _ := json.Marshal(accessory)
    assert.NotNil(t, result)
}