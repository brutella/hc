package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

type TemperatureSensor struct {
	*Service

	Temp *characteristic.TemperatureCharacteristic
	Name *characteristic.Name
}

// NewTemperatureSensor returns a temperature service service.
func NewTemperatureSensor(name string, temperature, min, max, steps float64) *TemperatureSensor {
	nameChar := characteristic.NewName(name)
	tempUnit := model.TempUnitCelsius
	temp := characteristic.NewCurrentTemperatureCharacteristic(temperature, min, max, steps, string(tempUnit))

	svc := New()
	svc.Type = typeTemperatureSensor
	svc.AddCharacteristic(temp.Characteristic)
	svc.AddCharacteristic(nameChar.Characteristic)

	t := TemperatureSensor{svc, temp, nameChar}

	return &t
}
