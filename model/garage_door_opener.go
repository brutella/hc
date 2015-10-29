package model

type DoorStateType byte

const (
	DoorStateOpen    DoorStateType = 0x00
	DoorStateClosed  DoorStateType = 0x01
	DoorStateOpening DoorStateType = 0x02
	DoorStateClosing DoorStateType = 0x03
    DoorStateStopped DoorStateType = 0x03
)

// A Garage Door Opener is a device that allows you to control a motor used
// to open a garage door.
//
// TODO(pmaene): Should be implemented as a lock.
// TODO(pmaene): The HAP protocol defines additional optional properties (lock current state
// and lock target state), which are not implemented yet.
type GarageDoorOpener interface {
	// SetState sets the current state
	SetState(DoorStateType)

	// State returns the current state
	State() DoorStateType

	// SetTargetState sets the target state
	SetTargetState(DoorStateType)

	// TargetState returns the target state
	TargetState() DoorStateType

    // SetObstruction sets the obstruction detection
    SetObstruction(bool)

    // Obstruction returns the obstruction detection
    Obstruction() bool
}
