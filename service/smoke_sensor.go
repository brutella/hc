// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSmokeSensor = "87"

type SmokeSensor struct {
	*Service

	SmokeDetected *characteristic.SmokeDetected

	StatusActive     *characteristic.StatusActive
	StatusFault      *characteristic.StatusFault
	StatusTampered   *characteristic.StatusTampered
	StatusLowBattery *characteristic.StatusLowBattery
	Name             *characteristic.Name
}

func NewSmokeSensor() *SmokeSensor {
	svc := SmokeSensor{}
	svc.Service = New(TypeSmokeSensor)

	svc.SmokeDetected = characteristic.NewSmokeDetected()
	svc.AddCharacteristic(svc.SmokeDetected.Characteristic)

	return &svc
}

func (svc *SmokeSensor) AddOptionalCharaterics() {

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
