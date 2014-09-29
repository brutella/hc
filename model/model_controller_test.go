package hk

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "bytes"
    "fmt"
    "io/ioutil"
)

func TestGetAccessories(t *testing.T) {
    info_service := NewAccessoryInfoService("123-456-789", "Rev1", "Matthias H.", "My Bridge")
    accessory := NewAccessory()
    accessory.AddService(info_service.Service)
    model := NewModel()
    model.AddAccessory(accessory)
    
    controller := NewModelController(model)
    
    var b bytes.Buffer
    r, err := controller.HandleGetAccessories(&b)
    assert.Nil(t, err)
    assert.NotNil(t, r)
    
    bytes, _ := ioutil.ReadAll(r)
    fmt.Println(string(bytes))
}