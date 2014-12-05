package accessory

import (
    "github.com/brutella/hap/model"
    
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestAccessory(t *testing.T) {    
    info := model.Info{
        Name: "My Accessory",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Accessory",
    }
    
    var a model.Accessory = New(info)
    
    assert.Equal(t, a.GetId(), model.InvalidId)
    assert.Equal(t, a.Name(), "My Accessory")
    assert.Equal(t, a.SerialNumber(), "001")
    assert.Equal(t, a.Manufacturer(), "Google")
    assert.Equal(t, a.Model(), "Accessory")
}