// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeHumidifierDehumidifier = "BD"

type HumidifierDehumidifier struct {
	*Service

	CurrentRelativeHumidity            *characteristic.CurrentRelativeHumidity
	CurrentHumidifierDehumidifierState *characteristic.CurrentHumidifierDehumidifierState
	TargetHumidifierDehumidifierState  *characteristic.TargetHumidifierDehumidifierState
	Active                             *characteristic.Active

	LockPhysicalControls                  *characteristic.LockPhysicalControls
	Name                                  *characteristic.Name
	SwingMode                             *characteristic.SwingMode
	WaterLevel                            *characteristic.WaterLevel
	RelativeHumidityDehumidifierThreshold *characteristic.RelativeHumidityDehumidifierThreshold
	RelativeHumidityHumidifierThreshold   *characteristic.RelativeHumidityHumidifierThreshold
	RotationSpeed                         *characteristic.RotationSpeed
}

func NewHumidifierDehumidifier() *HumidifierDehumidifier {
	svc := HumidifierDehumidifier{}
	svc.Service = New(TypeHumidifierDehumidifier)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddCharacteristic(svc.CurrentRelativeHumidity.Characteristic)

	svc.CurrentHumidifierDehumidifierState = characteristic.NewCurrentHumidifierDehumidifierState()
	svc.AddCharacteristic(svc.CurrentHumidifierDehumidifierState.Characteristic)

	svc.TargetHumidifierDehumidifierState = characteristic.NewTargetHumidifierDehumidifierState()
	svc.AddCharacteristic(svc.TargetHumidifierDehumidifierState.Characteristic)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	return &svc
}

func (svc *HumidifierDehumidifier) AddOptionalCharaterics() {

	svc.LockPhysicalControls = characteristic.NewLockPhysicalControls()
	svc.AddCharacteristic(svc.LockPhysicalControls.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.SwingMode = characteristic.NewSwingMode()
	svc.AddCharacteristic(svc.SwingMode.Characteristic)

	svc.WaterLevel = characteristic.NewWaterLevel()
	svc.AddCharacteristic(svc.WaterLevel.Characteristic)

	svc.RelativeHumidityDehumidifierThreshold = characteristic.NewRelativeHumidityDehumidifierThreshold()
	svc.AddCharacteristic(svc.RelativeHumidityDehumidifierThreshold.Characteristic)

	svc.RelativeHumidityHumidifierThreshold = characteristic.NewRelativeHumidityHumidifierThreshold()
	svc.AddCharacteristic(svc.RelativeHumidityHumidifierThreshold.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

}
