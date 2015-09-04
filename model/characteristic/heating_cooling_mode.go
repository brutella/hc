package characteristic

import (
	"github.com/brutella/hc/model"
)

type HeatingCoolingMode struct {
	*ByteCharacteristic
}

func NewHeatingCoolingMode(current model.HeatCoolModeType, CharacteristicType CharacteristicType, permissions []string) *HeatingCoolingMode {
	c := HeatingCoolingMode{NewByteCharacteristic(byte(current), permissions)}
	c.Type = CharacteristicType

	return &c
}

func NewCurrentHeatingCoolingMode(current model.HeatCoolModeType) *HeatingCoolingMode {
	return NewHeatingCoolingMode(current, TypeCurrentHeatingCoolingMode, PermsRead())
}

func NewTargetHeatingCoolingMode(current model.HeatCoolModeType) *HeatingCoolingMode {
	return NewHeatingCoolingMode(current, TypeTargetHeatingCoolingMode, PermsAll())
}

func (c *HeatingCoolingMode) SetHeatingCoolingMode(mode model.HeatCoolModeType) {
	c.SetByte(byte(mode))
}

func (c *HeatingCoolingMode) HeatingCoolingMode() model.HeatCoolModeType {
	return model.HeatCoolModeType(c.Byte())
}

// type CurrentRelativeHumidityCharacteristic struct {
//     *Float
//     humidity float64
// }
//
// func NewCurrentRelativeHumidityCharacteristic(value float64) *CurrentRelativeHumidityCharacteristic {
//     c := CurrentRelativeHumidityCharacteristic{NewFloat(value), value}
//     c.Type = TypeCurrentRelativeHumidity
//     c.Permissions = PermsRead()
//     return &c
// }
//
// type TargetRelativeHumidityCharacteristic struct {
//     *Float
//     target float64
// }
//
// func NewTargetRelativeHumidityCharacteristic(value, min, max, steps float64) *TargetRelativeHumidityCharacteristic {
//     c := TargetRelativeHumidityCharacteristic{NewFloatMinMaxSteps(value, min, max, steps), value}
//     c.Type = TypeTargetRelativeHumidity
//     c.Permissions = PermsAll()
//     return &c
// }
//
// type CoolingThresholdTemperatureCharacteristic struct {
//     *Float
// }
//
// func NewCoolingThresholdTemperatureCharacteristic(value, min, max, steps float64) *CoolingThresholdTemperatureCharacteristic {
//     c := CoolingThresholdTemperatureCharacteristic{NewFloatMinMaxSteps(value, min, max, steps), value}
//     c.Type = TypeCoolingThreshold
//     c.Permissions = PermsAll()
//     return &c
// }
//
// type HeatingThresholdTemperatureCharacteristic struct {
//     *Float
// }
//
// func NewHeatingThresholdTemperatureCharacteristic(value, min, max, steps float64) *HeatingThresholdTemperatureCharacteristic {
//     c := HeatingThresholdTemperatureCharacteristic{NewFloatMinMaxSteps(value, min, max, steps), value}
//     c.Type = TypeHeatingThresholdTemperature
//     c.Permissions = PermsAll()
//     return &c
// }
