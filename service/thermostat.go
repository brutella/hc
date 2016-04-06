// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeThermostat = "0000004A-0000-1000-8000-0026BB765291"

type Thermostat struct {
	*Service

	CurrentHeatingCoolingState *characteristic.CurrentHeatingCoolingState
	TargetHeatingCoolingState  *characteristic.TargetHeatingCoolingState
	CurrentTemperature         *characteristic.CurrentTemperature
	TargetTemperature          *characteristic.TargetTemperature
	TemperatureDisplayUnits    *characteristic.TemperatureDisplayUnits
}

func NewThermostat() *Thermostat {
	svc := Thermostat{}
	svc.Service = New(TypeThermostat)

	svc.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	svc.AddCharacteristic(svc.CurrentHeatingCoolingState.Characteristic)

	svc.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	svc.AddCharacteristic(svc.TargetHeatingCoolingState.Characteristic)

	svc.CurrentTemperature = characteristic.NewCurrentTemperature()
	svc.AddCharacteristic(svc.CurrentTemperature.Characteristic)

	svc.TargetTemperature = characteristic.NewTargetTemperature()
	svc.AddCharacteristic(svc.TargetTemperature.Characteristic)

	svc.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	svc.AddCharacteristic(svc.TemperatureDisplayUnits.Characteristic)

	return &svc
}
