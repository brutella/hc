package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

// Thermostat is service to represent a thermostat.
type Thermostat struct {
	*TemperatureSensor

	TargetTemp *characteristic.TemperatureCharacteristic
	Mode       *characteristic.HeatingCoolingMode
	TargetMode *characteristic.HeatingCoolingMode

	targetTempChange func(float64)
}

// NewThermostat returns a thermostat service.
func NewThermostat(name string, temperature, min, max, steps float64) *Thermostat {

	svc := NewTemperatureSensor(name, temperature, min, max, steps)

	tempUnit := svc.Unit.Unit()
	targetTemp := characteristic.NewTargetTemperatureCharacteristic(temperature, min, max, steps, string(tempUnit))
	mode := characteristic.NewCurrentHeatingCoolingMode(model.HeatCoolModeOff)
	targetMode := characteristic.NewTargetHeatingCoolingMode(model.HeatCoolModeOff)

	svc.Type = typeThermostat
	svc.AddCharacteristic(mode.Characteristic)
	svc.AddCharacteristic(targetMode.Characteristic)
	svc.AddCharacteristic(targetTemp.Characteristic)

	t := Thermostat{svc, targetTemp, mode, targetMode, nil}

	return &t
}
