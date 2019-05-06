// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeAirPurifier = "BB"

type AirPurifier struct {
	*Service

	Active                  *characteristic.Active
	CurrentAirPurifierState *characteristic.CurrentAirPurifierState
	TargetAirPurifierState  *characteristic.TargetAirPurifierState

	LockPhysicalControls *characteristic.LockPhysicalControls
	Name                 *characteristic.Name
	SwingMode            *characteristic.SwingMode
	RotationSpeed        *characteristic.RotationSpeed
}

func NewAirPurifier() *AirPurifier {
	svc := AirPurifier{}
	svc.Service = New(TypeAirPurifier)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentAirPurifierState = characteristic.NewCurrentAirPurifierState()
	svc.AddCharacteristic(svc.CurrentAirPurifierState.Characteristic)

	svc.TargetAirPurifierState = characteristic.NewTargetAirPurifierState()
	svc.AddCharacteristic(svc.TargetAirPurifierState.Characteristic)

	return &svc
}

func (svc *AirPurifier) addOptionalCharaterics() {

	svc.LockPhysicalControls = characteristic.NewLockPhysicalControls()
	svc.AddCharacteristic(svc.LockPhysicalControls.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.SwingMode = characteristic.NewSwingMode()
	svc.AddCharacteristic(svc.SwingMode.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

}
