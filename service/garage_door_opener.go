// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeGarageDoorOpener = "41"

type GarageDoorOpener struct {
	*Service

	CurrentDoorState    *characteristic.CurrentDoorState
	TargetDoorState     *characteristic.TargetDoorState
	ObstructionDetected *characteristic.ObstructionDetected

	LockCurrentState *characteristic.LockCurrentState
	LockTargetState  *characteristic.LockTargetState
	Name             *characteristic.Name
}

func NewGarageDoorOpener() *GarageDoorOpener {
	svc := GarageDoorOpener{}
	svc.Service = New(TypeGarageDoorOpener)

	svc.CurrentDoorState = characteristic.NewCurrentDoorState()
	svc.AddCharacteristic(svc.CurrentDoorState.Characteristic)

	svc.TargetDoorState = characteristic.NewTargetDoorState()
	svc.AddCharacteristic(svc.TargetDoorState.Characteristic)

	svc.ObstructionDetected = characteristic.NewObstructionDetected()
	svc.AddCharacteristic(svc.ObstructionDetected.Characteristic)

	svc.LockCurrentState = characteristic.NewLockCurrentState()
	svc.AddCharacteristic(svc.LockCurrentState.Characteristic)

	svc.LockTargetState = characteristic.NewLockTargetState()
	svc.AddCharacteristic(svc.LockTargetState.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}
