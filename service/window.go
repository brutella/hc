// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeWindow = "8B"

type Window struct {
	*Service

	CurrentPosition *characteristic.CurrentPosition
	TargetPosition  *characteristic.TargetPosition
	PositionState   *characteristic.PositionState
}

func NewWindow() *Window {
	svc := Window{}
	svc.Service = New(TypeWindow)

	svc.CurrentPosition = characteristic.NewCurrentPosition()
	svc.AddCharacteristic(svc.CurrentPosition.Characteristic)

	svc.TargetPosition = characteristic.NewTargetPosition()
	svc.AddCharacteristic(svc.TargetPosition.Characteristic)

	svc.PositionState = characteristic.NewPositionState()
	svc.AddCharacteristic(svc.PositionState.Characteristic)

	return &svc
}
