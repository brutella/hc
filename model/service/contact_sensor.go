package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

// ContactSensor is a service to represent a contact sensor.
type ContactSensor struct {
	*Service
	ContactSensorState *characteristic.ContactSensorState
	Name *characteristic.Name
}

// NewContactSensor returns a contact sensor service.
func NewContactSensor(name string) *ContactSensor {
	contactSensorChar := characteristic.NewCurrentContactSensorState(model.ContactNotDetected)
	nameChar := characteristic.NewName(name)

	svc := New()
	svc.Type = typeContactSensor
	svc.AddCharacteristic(contactSensorChar.Characteristic)
	svc.AddCharacteristic(nameChar.Characteristic)

	return &ContactSensor{svc, contactSensorChar, nameChar}
}
