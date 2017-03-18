package characteristic

import (
	"github.com/brutella/hc/model"
)

type ContactSensorState struct {
	*ByteCharacteristic
}

func NewContactSensorState(current model.ContactSensorStateType, CharacteristicType CharacteristicType, permissions []string) *ContactSensorState {
	c := ContactSensorState{NewByteCharacteristic(byte(current), permissions)}
	c.Type = CharacteristicType

	return &c
}

func NewCurrentContactSensorState(current model.ContactSensorStateType) *ContactSensorState {
	return NewContactSensorState(current, TypeContactSensorState, PermsRead())
}

func (c *ContactSensorState) SetContactSensorState(state model.ContactSensorStateType) {
	c.SetByte(byte(state))
}

func (c *ContactSensorState) ContactSensorState() model.ContactSensorStateType {
	return model.ContactSensorStateType(c.Byte())
}