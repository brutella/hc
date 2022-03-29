package accessory

import "github.com/brutella/hc/service"

type DoorLocks struct {
	*Accessory
	DoorLock *service.LockMechanism
}

// NewDoorLock returns a window which implements model.NewDoorLock.
func NewDoorLock(info Info, targetState int) *DoorLocks {
	acc := DoorLocks{}
	acc.Accessory = New(info, TypeDoorLock)
	acc.DoorLock = service.NewLockMechanism()
	acc.DoorLock.LockTargetState.SetValue(targetState)
	acc.AddService(acc.DoorLock.Service)

	return &acc
}

