// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeDoor = "81"

type Door struct {
	*Service

	CurrentPosition *characteristic.CurrentPosition
	PositionState   *characteristic.PositionState
	TargetPosition  *characteristic.TargetPosition

	HoldPosition        *characteristic.HoldPosition
	ObstructionDetected *characteristic.ObstructionDetected
	Name                *characteristic.Name
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

func (svc *Door) AddOptionalCharacteristics() {

	svc.HoldPosition = characteristic.NewHoldPosition()
	svc.AddCharacteristic(svc.HoldPosition.Characteristic)

	svc.ObstructionDetected = characteristic.NewObstructionDetected()
	svc.AddCharacteristic(svc.ObstructionDetected.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
