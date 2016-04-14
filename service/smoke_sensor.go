// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSmokeSensor = "87"

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
