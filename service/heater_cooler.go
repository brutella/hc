// THIS FILE IS AUTO-GENERATED
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

	LockPhysicalControls        *characteristic.LockPhysicalControls
	Name                        *characteristic.Name
	SwingMode                   *characteristic.SwingMode
	CoolingThresholdTemperature *characteristic.CoolingThresholdTemperature
	HeatingThresholdTemperature *characteristic.HeatingThresholdTemperature
	TemperatureDisplayUnits     *characteristic.TemperatureDisplayUnits
	RotationSpeed               *characteristic.RotationSpeed
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

	return &svc
}

func (svc *HeaterCooler) AddOptionalCharaterics() {

	svc.LockPhysicalControls = characteristic.NewLockPhysicalControls()
	svc.AddCharacteristic(svc.LockPhysicalControls.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.SwingMode = characteristic.NewSwingMode()
	svc.AddCharacteristic(svc.SwingMode.Characteristic)

	svc.CoolingThresholdTemperature = characteristic.NewCoolingThresholdTemperature()
	svc.AddCharacteristic(svc.CoolingThresholdTemperature.Characteristic)

	svc.HeatingThresholdTemperature = characteristic.NewHeatingThresholdTemperature()
	svc.AddCharacteristic(svc.HeatingThresholdTemperature.Characteristic)

	svc.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	svc.AddCharacteristic(svc.TemperatureDisplayUnits.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

}
