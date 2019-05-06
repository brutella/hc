// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLeakSensor = "83"

type LeakSensor struct {
	*Service

	LeakDetected *characteristic.LeakDetected

	StatusActive     *characteristic.StatusActive
	StatusFault      *characteristic.StatusFault
	StatusTampered   *characteristic.StatusTampered
	StatusLowBattery *characteristic.StatusLowBattery
	Name             *characteristic.Name
}

func NewLeakSensor() *LeakSensor {
	svc := LeakSensor{}
	svc.Service = New(TypeLeakSensor)

	svc.LeakDetected = characteristic.NewLeakDetected()
	svc.AddCharacteristic(svc.LeakDetected.Characteristic)

	return &svc
}

func (svc *LeakSensor) AddOptionalCharaterics() {

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
