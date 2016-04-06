// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLockManagement = "00000044-0000-1000-8000-0026BB765291"

type LockManagement struct {
	*Service

	LockControlPoint *characteristic.LockControlPoint
	Version          *characteristic.Version
}

func NewLockManagement() *LockManagement {
	svc := LockManagement{}
	svc.Service = New(TypeLockManagement)

	svc.LockControlPoint = characteristic.NewLockControlPoint()
	svc.AddCharacteristic(svc.LockControlPoint.Characteristic)

	svc.Version = characteristic.NewVersion()
	svc.AddCharacteristic(svc.Version.Characteristic)

	return &svc
}
