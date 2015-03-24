package service

import (
	"github.com/brutella/hc/model/characteristic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwitch(t *testing.T) {
	sw := NewSwitch("My Switch", true)

	assert.Equal(t, sw.Type, typeSwitch)
	assert.Equal(t, sw.Name.GetValue(), "My Switch")
	assert.Equal(t, sw.Name.Permissions, characteristic.PermsRead())
	assert.Equal(t, sw.On.GetValue(), true)
	assert.Equal(t, sw.On.Permissions, characteristic.PermsAll())
}
