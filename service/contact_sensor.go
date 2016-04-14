// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeContactSensor = "80"

type ContactSensor struct {
	*Service

	ContactSensorState *characteristic.ContactSensorState
}

func NewContactSensor() *ContactSensor {
	svc := ContactSensor{}
	svc.Service = New(TypeContactSensor)

	svc.ContactSensorState = characteristic.NewContactSensorState()
	svc.AddCharacteristic(svc.ContactSensorState.Characteristic)

	return &svc
}
