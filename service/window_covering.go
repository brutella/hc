// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeWindowCovering = "8C"

type WindowCovering struct {
	*Service

	CurrentPosition *characteristic.CurrentPosition
	TargetPosition  *characteristic.TargetPosition
	PositionState   *characteristic.PositionState

	HoldPosition               *characteristic.HoldPosition
	TargetHorizontalTiltAngle  *characteristic.TargetHorizontalTiltAngle
	TargetVerticalTiltAngle    *characteristic.TargetVerticalTiltAngle
	CurrentHorizontalTiltAngle *characteristic.CurrentHorizontalTiltAngle
	CurrentVerticalTiltAngle   *characteristic.CurrentVerticalTiltAngle
	ObstructionDetected        *characteristic.ObstructionDetected
	Name                       *characteristic.Name
}

func NewWindowCovering() *WindowCovering {
	svc := WindowCovering{}
	svc.Service = New(TypeWindowCovering)

	svc.CurrentPosition = characteristic.NewCurrentPosition()
	svc.AddCharacteristic(svc.CurrentPosition.Characteristic)

	svc.TargetPosition = characteristic.NewTargetPosition()
	svc.AddCharacteristic(svc.TargetPosition.Characteristic)

	svc.PositionState = characteristic.NewPositionState()
	svc.AddCharacteristic(svc.PositionState.Characteristic)

	return &svc
}

func (svc *WindowCovering) AddOptionalCharacteristics() {

	svc.HoldPosition = characteristic.NewHoldPosition()
	svc.AddCharacteristic(svc.HoldPosition.Characteristic)

	svc.TargetHorizontalTiltAngle = characteristic.NewTargetHorizontalTiltAngle()
	svc.AddCharacteristic(svc.TargetHorizontalTiltAngle.Characteristic)

	svc.TargetVerticalTiltAngle = characteristic.NewTargetVerticalTiltAngle()
	svc.AddCharacteristic(svc.TargetVerticalTiltAngle.Characteristic)

	svc.CurrentHorizontalTiltAngle = characteristic.NewCurrentHorizontalTiltAngle()
	svc.AddCharacteristic(svc.CurrentHorizontalTiltAngle.Characteristic)

	svc.CurrentVerticalTiltAngle = characteristic.NewCurrentVerticalTiltAngle()
	svc.AddCharacteristic(svc.CurrentVerticalTiltAngle.Characteristic)

	svc.ObstructionDetected = characteristic.NewObstructionDetected()
	svc.AddCharacteristic(svc.ObstructionDetected.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
