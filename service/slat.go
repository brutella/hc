// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSlat = "B9"

type Slat struct {
	*Service

	SlatType         *characteristic.SlatType
	CurrentSlatState *characteristic.CurrentSlatState

	Name             *characteristic.Name
	CurrentTiltAngle *characteristic.CurrentTiltAngle
	TargetTiltAngle  *characteristic.TargetTiltAngle
	SwingMode        *characteristic.SwingMode
}

func NewSlat() *Slat {
	svc := Slat{}
	svc.Service = New(TypeSlat)

	svc.SlatType = characteristic.NewSlatType()
	svc.AddCharacteristic(svc.SlatType.Characteristic)

	svc.CurrentSlatState = characteristic.NewCurrentSlatState()
	svc.AddCharacteristic(svc.CurrentSlatState.Characteristic)

	return &svc
}

func (svc *Slat) AddOptionalCharaterics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.CurrentTiltAngle = characteristic.NewCurrentTiltAngle()
	svc.AddCharacteristic(svc.CurrentTiltAngle.Characteristic)

	svc.TargetTiltAngle = characteristic.NewTargetTiltAngle()
	svc.AddCharacteristic(svc.TargetTiltAngle.Characteristic)

	svc.SwingMode = characteristic.NewSwingMode()
	svc.AddCharacteristic(svc.SwingMode.Characteristic)

}
