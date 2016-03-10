package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

type TemperatureSensor struct {
	*Service

	Temp *characteristic.TemperatureCharacteristic
	Name *characteristic.Name
	Unit *characteristic.TemperatureUnit
}

// NewTemperatureSensor returns a temperature service service.
func NewTemperatureSensor(name string, temperature, min, max, steps float64) *TemperatureSensor {
	nameChar := characteristic.NewName(name)
	tempUnit := model.TempUnitCelsius
	unitChar := characteristic.NewTemperatureUnit(tempUnit)
	temp := characteristic.NewCurrentTemperatureCharacteristic(temperature, min, max, steps, string(tempUnit))

	svc := New()
	svc.Type = TypeTemperatureSensor
	svc.AddCharacteristic(temp.Characteristic)
	svc.AddCharacteristic(nameChar.Characteristic)
	svc.AddCharacteristic(unitChar.Characteristic)

	t := TemperatureSensor{svc, temp, nameChar, unitChar}

	return &t
}
