package service

import (
	"github.com/brutella/hc/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThermostat(t *testing.T) {
	thermostat := NewThermostat("Testthermostat", 10.5, -10, 100, 1)

	assert.Equal(t, thermostat.Type, typeThermostat)
	assert.Equal(t, thermostat.Name.GetValue(), "Testthermostat")
	assert.Equal(t, thermostat.Temp.GetValue(), 10.5)
	assert.Equal(t, thermostat.TargetTemp.GetValue(), 10.5)
	assert.Equal(t, thermostat.Mode.GetValue(), model.HeatCoolModeOff)
	assert.Equal(t, thermostat.TargetMode.GetValue(), model.HeatCoolModeOff)
}

func TestThermometer(t *testing.T) {
	thermometer := NewThermometer("Thermometer", 10.5, -10, 100, 1)

	assert.Equal(t, thermometer.Type, typeThermostat)
	assert.Equal(t, thermometer.Name.GetValue(), "Thermometer")
	assert.Equal(t, thermometer.Temp.GetValue(), 10.5)
	assert.Equal(t, thermometer.TargetTemp.GetValue(), 10.5)
	assert.Equal(t, thermometer.Mode.GetValue(), model.HeatCoolModeOff)
	assert.Equal(t, thermometer.TargetMode.GetValue(), model.HeatCoolModeOff)
}
