// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeHumiditySensor = "82"

type HumiditySensor struct {
	*Service

	CurrentRelativeHumidity *characteristic.CurrentRelativeHumidity

	StatusActive     *characteristic.StatusActive
	StatusFault      *characteristic.StatusFault
	StatusTampered   *characteristic.StatusTampered
	StatusLowBattery *characteristic.StatusLowBattery
	Name             *characteristic.Name
}

func NewHumiditySensor() *HumiditySensor {
	svc := HumiditySensor{}
	svc.Service = New(TypeHumiditySensor)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddCharacteristic(svc.CurrentRelativeHumidity.Characteristic)

	return &svc
}

func (svc *HumiditySensor) AddOptionalCharaterics() {

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
