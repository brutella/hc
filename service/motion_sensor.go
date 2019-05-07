// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeMotionSensor = "85"

type MotionSensor struct {
	*Service

	MotionDetected *characteristic.MotionDetected

	StatusActive     *characteristic.StatusActive
	StatusFault      *characteristic.StatusFault
	StatusTampered   *characteristic.StatusTampered
	StatusLowBattery *characteristic.StatusLowBattery
	Name             *characteristic.Name
}

func NewMotionSensor() *MotionSensor {
	svc := MotionSensor{}
	svc.Service = New(TypeMotionSensor)

	svc.MotionDetected = characteristic.NewMotionDetected()
	svc.AddCharacteristic(svc.MotionDetected.Characteristic)

	return &svc
}

func (svc *MotionSensor) AddOptionalCharacteristics() {

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
