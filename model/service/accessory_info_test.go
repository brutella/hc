package service

import (
	"github.com/brutella/hc/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessoryInfo(t *testing.T) {
	info := model.Info{
		Name:         "Test Accessory",
		SerialNumber: "001",
		Manufacturer: "Matthias",
		Model:        "Version 123",
	}

	i := NewInfo(info)

	assert.Equal(t, i.Type, typeAccessoryInfo)
	assert.Nil(t, i.Identify.GetValue())
	assert.Equal(t, i.Serial.GetValue(), "001")
	assert.Equal(t, i.Model.GetValue(), "Version 123")
	assert.Equal(t, i.Manufacturer.GetValue(), "Matthias")
	assert.Equal(t, i.Name.GetValue(), "Test Accessory")
	assert.Nil(t, i.Firmware)
	assert.Nil(t, i.Hardware)
	assert.Nil(t, i.Software)
}
