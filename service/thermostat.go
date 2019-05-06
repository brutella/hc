// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeThermostat = "4A"

type Thermostat struct {
	*Service

	CurrentHeatingCoolingState *characteristic.CurrentHeatingCoolingState
	TargetHeatingCoolingState  *characteristic.TargetHeatingCoolingState
	CurrentTemperature         *characteristic.CurrentTemperature
	TargetTemperature          *characteristic.TargetTemperature
	TemperatureDisplayUnits    *characteristic.TemperatureDisplayUnits

	CurrentRelativeHumidity     *characteristic.CurrentRelativeHumidity
	TargetRelativeHumidity      *characteristic.TargetRelativeHumidity
	CoolingThresholdTemperature *characteristic.CoolingThresholdTemperature
	HeatingThresholdTemperature *characteristic.HeatingThresholdTemperature
	Name                        *characteristic.Name
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

func (svc *Thermostat) AddOptionalCharaterics() {

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddCharacteristic(svc.CurrentRelativeHumidity.Characteristic)

	svc.TargetRelativeHumidity = characteristic.NewTargetRelativeHumidity()
	svc.AddCharacteristic(svc.TargetRelativeHumidity.Characteristic)

	svc.CoolingThresholdTemperature = characteristic.NewCoolingThresholdTemperature()
	svc.AddCharacteristic(svc.CoolingThresholdTemperature.Characteristic)

	svc.HeatingThresholdTemperature = characteristic.NewHeatingThresholdTemperature()
	svc.AddCharacteristic(svc.HeatingThresholdTemperature.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
