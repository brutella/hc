// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCarbonMonoxideSensor = "7F"

type CarbonMonoxideSensor struct {
	*Service

	CarbonMonoxideDetected *characteristic.CarbonMonoxideDetected

	StatusActive            *characteristic.StatusActive
	StatusFault             *characteristic.StatusFault
	StatusLowBattery        *characteristic.StatusLowBattery
	StatusTampered          *characteristic.StatusTampered
	CarbonMonoxideLevel     *characteristic.CarbonMonoxideLevel
	CarbonMonoxidePeakLevel *characteristic.CarbonMonoxidePeakLevel
	Name                    *characteristic.Name
}

func NewCarbonMonoxideSensor() *CarbonMonoxideSensor {
	svc := CarbonMonoxideSensor{}
	svc.Service = New(TypeCarbonMonoxideSensor)

	svc.CarbonMonoxideDetected = characteristic.NewCarbonMonoxideDetected()
	svc.AddCharacteristic(svc.CarbonMonoxideDetected.Characteristic)

	return &svc
}

func (svc *CarbonMonoxideSensor) AddOptionalCharacteristics() {

	svc.StatusActive = characteristic.NewStatusActive()
	svc.AddCharacteristic(svc.StatusActive.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusLowBattery = characteristic.NewStatusLowBattery()
	svc.AddCharacteristic(svc.StatusLowBattery.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	svc.CarbonMonoxideLevel = characteristic.NewCarbonMonoxideLevel()
	svc.AddCharacteristic(svc.CarbonMonoxideLevel.Characteristic)

	svc.CarbonMonoxidePeakLevel = characteristic.NewCarbonMonoxidePeakLevel()
	svc.AddCharacteristic(svc.CarbonMonoxidePeakLevel.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
