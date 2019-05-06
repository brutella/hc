// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeOccupancySensor = "86"

type OccupancySensor struct {
	*Service

	OccupancyDetected *characteristic.OccupancyDetected

	Name             *characteristic.Name
	StatusActive     *characteristic.StatusActive
	StatusFault      *characteristic.StatusFault
	StatusTampered   *characteristic.StatusTampered
	StatusLowBattery *characteristic.StatusLowBattery
}

func NewOccupancySensor() *OccupancySensor {
	svc := OccupancySensor{}
	svc.Service = New(TypeOccupancySensor)

	svc.OccupancyDetected = characteristic.NewOccupancyDetected()
	svc.AddCharacteristic(svc.OccupancyDetected.Characteristic)

	return &svc
}

func (svc *OccupancySensor) addOptionalCharaterics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.StatusActive = characteristic.NewStatusActive()
	svc.AddCharacteristic(svc.StatusActive.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	svc.StatusLowBattery = characteristic.NewStatusLowBattery()
	svc.AddCharacteristic(svc.StatusLowBattery.Characteristic)

}
