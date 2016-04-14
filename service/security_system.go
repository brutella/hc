// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSecuritySystem = "7E"

type SecuritySystem struct {
	*Service

	SecuritySystemCurrentState *characteristic.SecuritySystemCurrentState
	SecuritySystemTargetState  *characteristic.SecuritySystemTargetState
}

func NewSecuritySystem() *SecuritySystem {
	svc := SecuritySystem{}
	svc.Service = New(TypeSecuritySystem)

	svc.SecuritySystemCurrentState = characteristic.NewSecuritySystemCurrentState()
	svc.AddCharacteristic(svc.SecuritySystemCurrentState.Characteristic)

	svc.SecuritySystemTargetState = characteristic.NewSecuritySystemTargetState()
	svc.AddCharacteristic(svc.SecuritySystemTargetState.Characteristic)

	return &svc
}
