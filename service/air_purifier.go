// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeAirPurifier = "BB"

type AirPurifier struct {
	*Service

	Active                  *characteristic.Active
	CurrentAirPurifierState *characteristic.CurrentAirPurifierState
	TargetAirPurifierState  *characteristic.TargetAirPurifierState
}

func NewAirPurifier() *AirPurifier {
	svc := AirPurifier{}
	svc.Service = New(TypeAirPurifier)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentAirPurifierState = characteristic.NewCurrentAirPurifierState()
	svc.AddCharacteristic(svc.CurrentAirPurifierState.Characteristic)

	svc.TargetAirPurifierState = characteristic.NewTargetAirPurifierState()
	svc.AddCharacteristic(svc.TargetAirPurifierState.Characteristic)

	return &svc
}
