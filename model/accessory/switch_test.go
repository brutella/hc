package accessory

import (
    "github.com/brutella/hap/model"
    
	"testing"
    "github.com/stretchr/testify/assert"
)
var info = model.Info{
    Name: "My Switch",
    SerialNumber: "001",
    Manufacturer: "Google",
    Model: "Switchy",
}

func TestSwitch(t *testing.T) {
    var s model.Switch = NewSwitch(info)
    
    assert.Equal(t, s.GetId(), model.InvalidId)
    assert.Equal(t, s.Name(), "My Switch")
    assert.Equal(t, s.SerialNumber(), "001")
    assert.Equal(t, s.Manufacturer(), "Google")
    assert.Equal(t, s.Model(), "Switchy")
    assert.False(t, s.IsOn())
    s.SetOn(true)
    assert.True(t, s.IsOn())
}

func TestSwitchOnChanged(t *testing.T) {
    s := NewSwitch(info)
    
    var newValue = false
    s.OnStateChanged(func(value bool) {
        newValue = value
    })
    
    s.switcher.On.SetValueFromRemote(true)
    assert.True(t, s.IsOn())
    assert.True(t, newValue)
}