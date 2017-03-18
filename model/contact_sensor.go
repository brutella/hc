package model

type ContactSensorStateType byte

const (
	ContactDetected    ContactSensorStateType = 0x00
	ContactNotDetected  ContactSensorStateType = 0x01
)

// A Contact Sensor is an accessory which provides contact state.
type ContactSensor interface {
	Accessory

	// SetState sets the current state
	SetState(ContactSensorStateType)

	// State returns the current state
	State() ContactSensorStateType
}