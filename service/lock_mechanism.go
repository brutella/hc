// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLockMechanism = "00000045-0000-1000-8000-0026BB765291"

type LockMechanism struct {
	*Service

	LockCurrentState *characteristic.LockCurrentState
	LockTargetState  *characteristic.LockTargetState
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
