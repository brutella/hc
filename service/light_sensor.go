// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLightSensor = "84"

type LightSensor struct {
	*Service

	CurrentAmbientLightLevel *characteristic.CurrentAmbientLightLevel

	Name             *characteristic.Name
	StatusActive     *characteristic.StatusActive
	StatusFault      *characteristic.StatusFault
	StatusTampered   *characteristic.StatusTampered
	StatusLowBattery *characteristic.StatusLowBattery
}

func NewLightSensor() *LightSensor {
	svc := LightSensor{}
	svc.Service = New(TypeLightSensor)

	svc.CurrentAmbientLightLevel = characteristic.NewCurrentAmbientLightLevel()
	svc.AddCharacteristic(svc.CurrentAmbientLightLevel.Characteristic)

	return &svc
}

func (svc *LightSensor) AddOptionalCharaterics() {

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
