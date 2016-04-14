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

	return &svc
}
