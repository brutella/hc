package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

// Thermostat is service to represent a thermostat.
type Thermostat struct {
	*Service

	Name       *characteristic.Name
	Unit       *characteristic.TemperatureUnit
	Temp       *characteristic.TemperatureCharacteristic
	TargetTemp *characteristic.TemperatureCharacteristic
	Mode       *characteristic.HeatingCoolingMode
	TargetMode *characteristic.HeatingCoolingMode

	targetTempChange func(float64)
}

// NewThermometer returns a thermometer service.
func NewThermometer(name string, temperature, min, max, steps float64) *Thermostat {
	thermostat := NewThermostat(name, temperature, min, max, steps)

	thermostat.TargetTemp.Permissions = characteristic.PermsRead()
	thermostat.TargetMode.Permissions = characteristic.PermsRead()

	return thermostat
}

// NewThermostat returns a thermostat service.
func NewThermostat(name string, temperature, min, max, steps float64) *Thermostat {
	nameChar := characteristic.NewName(name)
	tempUnit := model.TempUnitCelsius
	unitChar := characteristic.NewTemperatureUnit(tempUnit)
	temp := characteristic.NewCurrentTemperatureCharacteristic(temperature, min, max, steps, string(tempUnit))
	targetTemp := characteristic.NewTargetTemperatureCharacteristic(temperature, min, max, steps, string(tempUnit))
	mode := characteristic.NewCurrentHeatingCoolingMode(model.HeatCoolModeOff)
	targetMode := characteristic.NewTargetHeatingCoolingMode(model.HeatCoolModeOff)

	service := New()
	service.Type = typeThermostat
	service.addCharacteristic(mode.Characteristic)
	service.addCharacteristic(targetMode.Characteristic)
	service.addCharacteristic(temp.Characteristic)
	service.addCharacteristic(targetTemp.Characteristic)
	service.addCharacteristic(unitChar.Characteristic)
	service.addCharacteristic(nameChar.Characteristic)

	t := Thermostat{service, nameChar, unitChar, temp, targetTemp, mode, targetMode, nil}

	return &t
}
