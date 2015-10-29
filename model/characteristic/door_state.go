package characteristic

import (
	"github.com/brutella/hc/model"
)

type DoorState struct {
	*ByteCharacteristic
}

func NewDoorState(current model.DoorStateType, CharacteristicType CharacteristicType, permissions []string) *DoorState {
	c := DoorState{NewByteCharacteristic(byte(current), permissions)}
	c.Type = CharacteristicType

	return &c
}

func NewCurrentDoorState(current model.DoorStateType) *DoorState {
	return NewHeatingCoolingMode(current, TypeCurrentDoorState, PermsRead())
}

func NewTargetDoorState(current model.DoorStateType) *DoorState {
	return NewHeatingCoolingMode(current, TypeTargetDoorState, PermsAll())
}

func (c *DoorState) SetDoorState(mode model.DoorStateType) {
	c.SetByte(byte(mode))
}

func (c *DoorState) DoorState() model.DoorStateType {
	return model.HeatCoolModeType(c.Byte())
}
