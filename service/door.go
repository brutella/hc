// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeDoor = "00000081-0000-1000-8000-0026BB765291"

type Door struct {
	*Service

	CurrentPosition *characteristic.CurrentPosition
	PositionState   *characteristic.PositionState
	TargetPosition  *characteristic.TargetPosition
}

func NewDoor() *Door {
	svc := Door{}
	svc.Service = New(TypeDoor)

	svc.CurrentPosition = characteristic.NewCurrentPosition()
	svc.AddCharacteristic(svc.CurrentPosition.Characteristic)

	svc.PositionState = characteristic.NewPositionState()
	svc.AddCharacteristic(svc.PositionState.Characteristic)

	svc.TargetPosition = characteristic.NewTargetPosition()
	svc.AddCharacteristic(svc.TargetPosition.Characteristic)

	return &svc
}
