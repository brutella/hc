package accessory

import (
	"github.com/brutella/hc/model"

	"github.com/stretchr/testify/assert"
	"testing"
)

var outlet_info = model.Info{
	Name:         "My Outlet",
	SerialNumber: "001",
	Manufacturer: "brutella",
	Model:        "Outletty",
}

func TestOutlet(t *testing.T) {
	var o model.Outlet = NewOutlet(outlet_info)

	assert.Equal(t, o.GetID(), model.InvalidID)
	assert.Equal(t, o.Name(), "My Outlet")
	assert.Equal(t, o.SerialNumber(), "001")
	assert.Equal(t, o.Manufacturer(), "brutella")
	assert.Equal(t, o.Model(), "Outletty")
	assert.False(t, o.IsOn())
	assert.False(t, o.IsInUse())
	o.SetOn(true)
	o.SetInUse(true)
	assert.True(t, o.IsOn())
	assert.True(t, o.IsInUse())
}

func TestOutletOnChanged(t *testing.T) {
	o := NewOutlet(outlet_info)

	var newValue = false
	o.OnStateChanged(func(value bool) {
		newValue = value
	})

	o.outlet.On.SetValueFromRemote(true)
	assert.True(t, o.IsOn())
	assert.True(t, newValue)
}

func TestOutletInUseChanged(t *testing.T) {
	o := NewOutlet(outlet_info)

	var newValue = false
	o.InUseStateChanged(func(value bool) {
		newValue = value
	})

	o.outlet.InUse.SetValueFromRemote(true)
	assert.True(t, o.IsInUse())
	assert.True(t, newValue)
}
