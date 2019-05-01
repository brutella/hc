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

	StatusFault             *characteristic.StatusFault
	StatusTampered          *characteristic.StatusTampered
	SecuritySystemAlarmType *characteristic.SecuritySystemAlarmType
	Name                    *characteristic.Name
}

func NewSecuritySystem() *SecuritySystem {
	svc := SecuritySystem{}
	svc.Service = New(TypeSecuritySystem)

	svc.SecuritySystemCurrentState = characteristic.NewSecuritySystemCurrentState()
	svc.AddCharacteristic(svc.SecuritySystemCurrentState.Characteristic)

	svc.SecuritySystemTargetState = characteristic.NewSecuritySystemTargetState()
	svc.AddCharacteristic(svc.SecuritySystemTargetState.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	svc.SecuritySystemAlarmType = characteristic.NewSecuritySystemAlarmType()
	svc.AddCharacteristic(svc.SecuritySystemAlarmType.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}
