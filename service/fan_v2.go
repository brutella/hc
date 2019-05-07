// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeFanV2 = "B7"

type FanV2 struct {
	*Service

	Active *characteristic.Active

	CurrentFanState      *characteristic.CurrentFanState
	TargetFanState       *characteristic.TargetFanState
	LockPhysicalControls *characteristic.LockPhysicalControls
	Name                 *characteristic.Name
	RotationDirection    *characteristic.RotationDirection
	RotationSpeed        *characteristic.RotationSpeed
	SwingMode            *characteristic.SwingMode
}

func NewFanV2() *FanV2 {
	svc := FanV2{}
	svc.Service = New(TypeFanV2)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	return &svc
}

func (svc *FanV2) AddOptionalCharacteristics() {

	svc.CurrentFanState = characteristic.NewCurrentFanState()
	svc.AddCharacteristic(svc.CurrentFanState.Characteristic)

	svc.TargetFanState = characteristic.NewTargetFanState()
	svc.AddCharacteristic(svc.TargetFanState.Characteristic)

	svc.LockPhysicalControls = characteristic.NewLockPhysicalControls()
	svc.AddCharacteristic(svc.LockPhysicalControls.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.RotationDirection = characteristic.NewRotationDirection()
	svc.AddCharacteristic(svc.RotationDirection.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	svc.SwingMode = characteristic.NewSwingMode()
	svc.AddCharacteristic(svc.SwingMode.Characteristic)

}
