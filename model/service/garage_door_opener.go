package service

import (
    "github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

// GarageDoorOpener is service to represent a garage door opener.
type GarageDoorOpener struct {
    *Service
	State               *characteristic.DoorState
    TargetState         *characteristic.DoorState
    ObstructionDetected *characteristic.ObstructionDetected
}

// NewGarageDoorOpener returns a garage door opener service.
func NewGarageDoorOpener(name string) *GarageDoorOpener {
	nameChar := characteristic.NewName(name)
    state := characteristic.NewCurrentDoorState(model.DoorStateClosed)
    targetState := characteristic.NewTargetDoorState(model.DoorStateClosed)
    obstructionDetected := characteristic.NewObstructionDetected(false)

    svc := New()
	svc.Type = typeGarageDoorOpener
	svc.AddCharacteristic(nameChar.Characteristic)
	svc.AddCharacteristic(state.Characteristic)
	svc.AddCharacteristic(targetState.Characteristic)
	svc.AddCharacteristic(hue.Characteristic)

	return &GarageDoorOpener{svc, state, targetState, obstructionDetected}
}
