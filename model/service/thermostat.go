package service

import (
	"github.com/brutella/hap/model"
	"github.com/brutella/hap/model/characteristic"
)

type TempChangeFunc func(float64)
type Thermostat struct {
	*Service

	Name       *characteristic.Name
	Unit       *characteristic.TemperatureUnit
	Temp       *characteristic.TemperatureCharacteristic
	TargetTemp *characteristic.TemperatureCharacteristic
	Mode       *characteristic.HeatingCoolingMode
	TargetMode *characteristic.HeatingCoolingMode

	targetTempChange TempChangeFunc
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
	name_char := characteristic.NewName(name)
	unit := model.TempUnitCelsius
	unit_char := characteristic.NewTemperatureUnit(unit)
	temp := characteristic.NewCurrentTemperatureCharacteristic(temperature, min, max, steps, string(unit))
	targetTemp := characteristic.NewTargetTemperatureCharacteristic(temperature, min, max, steps, string(unit))
	mode := characteristic.NewCurrentHeatingCoolingMode(model.ModeOff)
	targetMode := characteristic.NewTargetHeatingCoolingMode(model.ModeOff)

	service := New()
	service.Type = TypeThermostat
	service.AddCharacteristic(name_char.Characteristic)
	service.AddCharacteristic(unit_char.Characteristic)
	service.AddCharacteristic(temp.Characteristic)
	service.AddCharacteristic(targetTemp.Characteristic)
	service.AddCharacteristic(mode.Characteristic)
	service.AddCharacteristic(targetMode.Characteristic)

	t := Thermostat{service, name_char, unit_char, temp, targetTemp, mode, targetMode, nil}

	return &t
}
