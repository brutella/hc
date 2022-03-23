// THIS FILE IS AUTO-GENERATED
// Modified by mestafin
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeHeaterCooler = "BC"

type HeaterCooler struct {
	*Service

	Active                   *characteristic.Active
	CurrentHeaterCoolerState *characteristic.CurrentHeaterCoolerState
	TargetHeaterCoolerState  *characteristic.TargetHeaterCoolerState
	CurrentTemperature       *characteristic.CurrentTemperature

	TargetTemperature *characteristic.TargetTemperature
}

func NewHeaterCooler() *HeaterCooler {
	svc := HeaterCooler{}
	svc.Service = New(TypeHeaterCooler)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentHeaterCoolerState = characteristic.NewCurrentHeaterCoolerState()
	svc.AddCharacteristic(svc.CurrentHeaterCoolerState.Characteristic)

	svc.TargetHeaterCoolerState = characteristic.NewTargetHeaterCoolerState()
	svc.AddCharacteristic(svc.TargetHeaterCoolerState.Characteristic)

	svc.CurrentTemperature = characteristic.NewCurrentTemperature()
	svc.AddCharacteristic(svc.CurrentTemperature.Characteristic)

	svc.TargetTemperature = characteristic.NewTargetTemperature()
	svc.AddCharacteristic(svc.TargetTemperature.Characteristic)

	return &svc
}
