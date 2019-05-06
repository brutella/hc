// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeTelevision = "D8"

type Television struct {
	*Service

	Active             *characteristic.Active
	ActiveIdentifier   *characteristic.ActiveIdentifier
	ConfiguredName     *characteristic.ConfiguredName
	SleepDiscoveryMode *characteristic.SleepDiscoveryMode

	Brightness         *characteristic.Brightness
	ClosedCaptions     *characteristic.ClosedCaptions
	DisplayOrder       *characteristic.DisplayOrder
	CurrentMediaState  *characteristic.CurrentMediaState
	TargetMediaState   *characteristic.TargetMediaState
	PictureMode        *characteristic.PictureMode
	PowerModeSelection *characteristic.PowerModeSelection
	RemoteKey          *characteristic.RemoteKey
}

func NewTelevision() *Television {
	svc := Television{}
	svc.Service = New(TypeTelevision)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.ActiveIdentifier = characteristic.NewActiveIdentifier()
	svc.AddCharacteristic(svc.ActiveIdentifier.Characteristic)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddCharacteristic(svc.ConfiguredName.Characteristic)

	svc.SleepDiscoveryMode = characteristic.NewSleepDiscoveryMode()
	svc.AddCharacteristic(svc.SleepDiscoveryMode.Characteristic)

	return &svc
}

func (svc *Television) addOptionalCharaterics() {

	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	svc.ClosedCaptions = characteristic.NewClosedCaptions()
	svc.AddCharacteristic(svc.ClosedCaptions.Characteristic)

	svc.DisplayOrder = characteristic.NewDisplayOrder()
	svc.AddCharacteristic(svc.DisplayOrder.Characteristic)

	svc.CurrentMediaState = characteristic.NewCurrentMediaState()
	svc.AddCharacteristic(svc.CurrentMediaState.Characteristic)

	svc.TargetMediaState = characteristic.NewTargetMediaState()
	svc.AddCharacteristic(svc.TargetMediaState.Characteristic)

	svc.PictureMode = characteristic.NewPictureMode()
	svc.AddCharacteristic(svc.PictureMode.Characteristic)

	svc.PowerModeSelection = characteristic.NewPowerModeSelection()
	svc.AddCharacteristic(svc.PowerModeSelection.Characteristic)

	svc.RemoteKey = characteristic.NewRemoteKey()
	svc.AddCharacteristic(svc.RemoteKey.Characteristic)

}
