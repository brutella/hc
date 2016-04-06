// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeTemperatureSensor = "0000008A-0000-1000-8000-0026BB765291"

type TemperatureSensor struct {
	*Service

	CurrentTemperature *characteristic.CurrentTemperature
}

func NewTemperatureSensor() *TemperatureSensor {
	svc := TemperatureSensor{}
	svc.Service = New(TypeTemperatureSensor)

	svc.CurrentTemperature = characteristic.NewCurrentTemperature()
	svc.AddCharacteristic(svc.CurrentTemperature.Characteristic)

	return &svc
}
