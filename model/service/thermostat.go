package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

// Thermostat is service to represent a thermostat.
type Thermostat struct {
	*TemperatureSensor

	Unit       *characteristic.TemperatureUnit
	TargetTemp *characteristic.TemperatureCharacteristic
	Mode       *characteristic.HeatingCoolingMode
	TargetMode *characteristic.HeatingCoolingMode

	targetTempChange func(float64)
}

// NewThermostat returns a thermostat service.
func NewThermostat(name string, temperature, min, max, steps float64) *Thermostat {
	tempUnit := model.TempUnitCelsius
	unitChar := characteristic.NewTemperatureUnit(tempUnit)
	targetTemp := characteristic.NewTargetTemperatureCharacteristic(temperature, min, max, steps, string(tempUnit))
	mode := characteristic.NewCurrentHeatingCoolingMode(model.HeatCoolModeOff)
	targetMode := characteristic.NewTargetHeatingCoolingMode(model.HeatCoolModeOff)

	svc := NewTemperatureSensor(name, temperature, min, max, steps)
	svc.Type = typeThermostat
	svc.AddCharacteristic(mode.Characteristic)
	svc.AddCharacteristic(targetMode.Characteristic)
	svc.AddCharacteristic(targetTemp.Characteristic)
	svc.AddCharacteristic(unitChar.Characteristic)

	t := Thermostat{svc, unitChar, targetTemp, mode, targetMode, nil}

	return &t
}
