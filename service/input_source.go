// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeInputSource = "D9"

type InputSource struct {
	*Service

	ConfiguredName         *characteristic.ConfiguredName
	InputSourceType        *characteristic.InputSourceType
	IsConfigured           *characteristic.IsConfigured
	CurrentVisibilityState *characteristic.CurrentVisibilityState

	Identifier            *characteristic.Identifier
	InputDeviceType       *characteristic.InputDeviceType
	TargetVisibilityState *characteristic.TargetVisibilityState
	Name                  *characteristic.Name
}

func NewInputSource() *InputSource {
	svc := InputSource{}
	svc.Service = New(TypeInputSource)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddCharacteristic(svc.ConfiguredName.Characteristic)

	svc.InputSourceType = characteristic.NewInputSourceType()
	svc.AddCharacteristic(svc.InputSourceType.Characteristic)

	svc.IsConfigured = characteristic.NewIsConfigured()
	svc.AddCharacteristic(svc.IsConfigured.Characteristic)

	svc.CurrentVisibilityState = characteristic.NewCurrentVisibilityState()
	svc.AddCharacteristic(svc.CurrentVisibilityState.Characteristic)

	svc.Identifier = characteristic.NewIdentifier()
	svc.AddCharacteristic(svc.Identifier.Characteristic)

	svc.InputDeviceType = characteristic.NewInputDeviceType()
	svc.AddCharacteristic(svc.InputDeviceType.Characteristic)

	svc.TargetVisibilityState = characteristic.NewTargetVisibilityState()
	svc.AddCharacteristic(svc.TargetVisibilityState.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}
