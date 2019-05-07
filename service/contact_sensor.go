// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeContactSensor = "80"

type ContactSensor struct {
	*Service

	ContactSensorState *characteristic.ContactSensorState

	StatusActive     *characteristic.StatusActive
	StatusFault      *characteristic.StatusFault
	StatusTampered   *characteristic.StatusTampered
	StatusLowBattery *characteristic.StatusLowBattery
	Name             *characteristic.Name
}

func NewContactSensor() *ContactSensor {
	svc := ContactSensor{}
	svc.Service = New(TypeContactSensor)

	svc.ContactSensorState = characteristic.NewContactSensorState()
	svc.AddCharacteristic(svc.ContactSensorState.Characteristic)

	return &svc
}

func (svc *ContactSensor) AddOptionalCharacteristics() {

	svc.StatusActive = characteristic.NewStatusActive()
	svc.AddCharacteristic(svc.StatusActive.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	svc.StatusLowBattery = characteristic.NewStatusLowBattery()
	svc.AddCharacteristic(svc.StatusLowBattery.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
