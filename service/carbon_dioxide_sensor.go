// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCarbonDioxideSensor = "97"

type CarbonDioxideSensor struct {
	*Service

	CarbonDioxideDetected *characteristic.CarbonDioxideDetected

	StatusActive           *characteristic.StatusActive
	StatusFault            *characteristic.StatusFault
	StatusLowBattery       *characteristic.StatusLowBattery
	StatusTampered         *characteristic.StatusTampered
	CarbonDioxideLevel     *characteristic.CarbonDioxideLevel
	CarbonDioxidePeakLevel *characteristic.CarbonDioxidePeakLevel
	Name                   *characteristic.Name
}

func NewCarbonDioxideSensor() *CarbonDioxideSensor {
	svc := CarbonDioxideSensor{}
	svc.Service = New(TypeCarbonDioxideSensor)

	svc.CarbonDioxideDetected = characteristic.NewCarbonDioxideDetected()
	svc.AddCharacteristic(svc.CarbonDioxideDetected.Characteristic)

	return &svc
}

func (svc *CarbonDioxideSensor) AddOptionalCharaterics() {

	svc.StatusActive = characteristic.NewStatusActive()
	svc.AddCharacteristic(svc.StatusActive.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusLowBattery = characteristic.NewStatusLowBattery()
	svc.AddCharacteristic(svc.StatusLowBattery.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	svc.CarbonDioxideLevel = characteristic.NewCarbonDioxideLevel()
	svc.AddCharacteristic(svc.CarbonDioxideLevel.Characteristic)

	svc.CarbonDioxidePeakLevel = characteristic.NewCarbonDioxidePeakLevel()
	svc.AddCharacteristic(svc.CarbonDioxidePeakLevel.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
