// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeWindowCovering = "0000008C-0000-1000-8000-0026BB765291"

type WindowCovering struct {
	*Service

	CurrentPosition *characteristic.CurrentPosition
	TargetPosition  *characteristic.TargetPosition
	PositionState   *characteristic.PositionState
}

func NewWindowCovering() *WindowCovering {
	svc := WindowCovering{}
	svc.Service = New(TypeWindowCovering)

	svc.CurrentPosition = characteristic.NewCurrentPosition()
	svc.AddCharacteristic(svc.CurrentPosition.Characteristic)

	svc.TargetPosition = characteristic.NewTargetPosition()
	svc.AddCharacteristic(svc.TargetPosition.Characteristic)

	svc.PositionState = characteristic.NewPositionState()
	svc.AddCharacteristic(svc.PositionState.Characteristic)

	return &svc
}
