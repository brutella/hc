// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeBatteryService = "96"

type BatteryService struct {
	*Service

	BatteryLevel     *characteristic.BatteryLevel
	ChargingState    *characteristic.ChargingState
	StatusLowBattery *characteristic.StatusLowBattery

	Name *characteristic.Name
}

func NewBatteryService() *BatteryService {
	svc := BatteryService{}
	svc.Service = New(TypeBatteryService)

	svc.BatteryLevel = characteristic.NewBatteryLevel()
	svc.AddCharacteristic(svc.BatteryLevel.Characteristic)

	svc.ChargingState = characteristic.NewChargingState()
	svc.AddCharacteristic(svc.ChargingState.Characteristic)

	svc.StatusLowBattery = characteristic.NewStatusLowBattery()
	svc.AddCharacteristic(svc.StatusLowBattery.Characteristic)

	return &svc
}

func (svc *BatteryService) addOptionalCharaterics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
