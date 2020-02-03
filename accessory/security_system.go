package accessory

import (
	"github.com/brutella/hc/service"
)

//SecuritySystem structure
type SecuritySystem struct {
	*Accessory
	SecuritySystem *service.SecuritySystem
}

// NewSecuritySystem returns an SecuritySystem accessory containing one SecuritySystem service.
func NewSecuritySystem(info Info, setSSTS, minSSTS, maxSSTS, stepSSTS int) *SecuritySystem {
	acc := SecuritySystem{}
	acc.Accessory = New(info, TypeSecuritySystem)
	acc.SecuritySystem = service.NewSecuritySystem()

	acc.SecuritySystem.SecuritySystemTargetState.SetValue(setSSTS)
	acc.SecuritySystem.SecuritySystemTargetState.SetMinValue(minSSTS)
	acc.SecuritySystem.SecuritySystemTargetState.SetMaxValue(maxSSTS)
	acc.SecuritySystem.SecuritySystemTargetState.SetStepValue(stepSSTS)

	acc.AddService(acc.SecuritySystem.Service)

	return &acc
}
