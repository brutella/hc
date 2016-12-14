// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeHumidifierDehumidifier = "BD"

type HumidifierDehumidifier struct {
	*Service

	CurrentRelativeHumidity            *characteristic.CurrentRelativeHumidity
	CurrentHumidifierDehumidifierState *characteristic.CurrentHumidifierDehumidifierState
	TargetHumidifierDehumidifierState  *characteristic.TargetHumidifierDehumidifierState
	Active                             *characteristic.Active
}

func NewHumidifierDehumidifier() *HumidifierDehumidifier {
	svc := HumidifierDehumidifier{}
	svc.Service = New(TypeHumidifierDehumidifier)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddCharacteristic(svc.CurrentRelativeHumidity.Characteristic)

	svc.CurrentHumidifierDehumidifierState = characteristic.NewCurrentHumidifierDehumidifierState()
	svc.AddCharacteristic(svc.CurrentHumidifierDehumidifierState.Characteristic)

	svc.TargetHumidifierDehumidifierState = characteristic.NewTargetHumidifierDehumidifierState()
	svc.AddCharacteristic(svc.TargetHumidifierDehumidifierState.Characteristic)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	return &svc
}
