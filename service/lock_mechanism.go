// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLockMechanism = "45"

type LockMechanism struct {
	*Service

	LockCurrentState *characteristic.LockCurrentState
	LockTargetState  *characteristic.LockTargetState

	Name *characteristic.Name
}

func NewLockMechanism() *LockMechanism {
	svc := LockMechanism{}
	svc.Service = New(TypeLockMechanism)

	svc.LockCurrentState = characteristic.NewLockCurrentState()
	svc.AddCharacteristic(svc.LockCurrentState.Characteristic)

	svc.LockTargetState = characteristic.NewLockTargetState()
	svc.AddCharacteristic(svc.LockTargetState.Characteristic)

	return &svc
}

func (svc *LockMechanism) AddOptionalCharacteristics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
