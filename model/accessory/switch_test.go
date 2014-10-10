package accessory

import (
    "github.com/brutella/hap/model"
    
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestSwitch(t *testing.T) {
    info := model.Info{
        Name: "My Switch",
        Serial: "001",
        Manufacturer: "Google",
        Model: "Switchy",
    }
    
    var s model.Switch = NewSwitch(info)
    
    assert.Equal(t, s.Name(), "My Switch")
    assert.Equal(t, s.SerialNumber(), "001")
    assert.Equal(t, s.Manufacturer(), "Google")
    assert.Equal(t, s.Model(), "Switchy")
    assert.False(t, s.IsOn())
    
    s.SetOn(true)
    
    assert.True(t, s.IsOn())
}