// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeFilterMaintenance = "BA"

type FilterMaintenance struct {
	*Service

	FilterChangeIndication *characteristic.FilterChangeIndication

	FilterLifeLevel       *characteristic.FilterLifeLevel
	ResetFilterIndication *characteristic.ResetFilterIndication
	Name                  *characteristic.Name
}

func NewFilterMaintenance() *FilterMaintenance {
	svc := FilterMaintenance{}
	svc.Service = New(TypeFilterMaintenance)

	svc.FilterChangeIndication = characteristic.NewFilterChangeIndication()
	svc.AddCharacteristic(svc.FilterChangeIndication.Characteristic)

	return &svc
}

func (svc *FilterMaintenance) addOptionalCharaterics() {

	svc.FilterLifeLevel = characteristic.NewFilterLifeLevel()
	svc.AddCharacteristic(svc.FilterLifeLevel.Characteristic)

	svc.ResetFilterIndication = characteristic.NewResetFilterIndication()
	svc.AddCharacteristic(svc.ResetFilterIndication.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
