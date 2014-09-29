package model

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "encoding/json"
    "fmt"
)

func TestModel(t *testing.T) {
    // serialNumber, modelName, manufacturerName, accessoryName string
    info_service := NewAccessoryInfoService("123-456-789", "Rev1", "Matthias H.", "My Bridge")
    accessory := NewAccessory()
    accessory.AddService(info_service.Service)
    
    model := NewModel()
    model.AddAccessory(accessory)
    result, _ := json.Marshal(model)
    assert.NotNil(t, result)
    fmt.Println(string(result))
}