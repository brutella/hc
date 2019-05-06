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

	HoldPosition        *characteristic.HoldPosition
	ObstructionDetected *characteristic.ObstructionDetected
	Name                *characteristic.Name
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

func (svc *Window) AddOptionalCharaterics() {

	svc.HoldPosition = characteristic.NewHoldPosition()
	svc.AddCharacteristic(svc.HoldPosition.Characteristic)

	svc.ObstructionDetected = characteristic.NewObstructionDetected()
	svc.AddCharacteristic(svc.ObstructionDetected.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
