// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSmokeSensor = "00000087-0000-1000-8000-0026BB765291"

type SmokeSensor struct {
	*Service

	SmokeDetected *characteristic.SmokeDetected
}

func NewSmokeSensor() *SmokeSensor {
	svc := SmokeSensor{}
	svc.Service = New(TypeSmokeSensor)

	svc.SmokeDetected = characteristic.NewSmokeDetected()
	svc.AddCharacteristic(svc.SmokeDetected.Characteristic)

	return &svc
}
