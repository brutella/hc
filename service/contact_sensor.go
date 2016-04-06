// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeContactSensor = "00000080-0000-1000-8000-0026BB765291"

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
