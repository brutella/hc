package characteristic

import(
    "github.com/brutella/hap/model"
)

type HeatingCoolingMode struct {
    *ByteCharacteristic
}

func NewHeatingCoolingMode(current model.HeatCoolMode, charType CharType, permissions []string) *HeatingCoolingMode {
    c := HeatingCoolingMode{NewByteCharacteristic(byte(current))}
    c.Type = charType
    c.Permissions = permissions
    return &c
}

func NewCurrentHeatingCoolingMode(current model.HeatCoolMode) *HeatingCoolingMode {
    return NewHeatingCoolingMode(current, CharTypeHeatingCoolingModeCurrent, PermsRead())
}

func NewTargetHeatingCoolingMode(current model.HeatCoolMode) *HeatingCoolingMode {
    return NewHeatingCoolingMode(current, CharTypeHeatingCoolingModeTarget, PermsRead())
}

func (c *HeatingCoolingMode) SetHeatingCoolingMode(mode model.HeatCoolMode) {
    c.SetByte(byte(mode))
}


func (c *HeatingCoolingMode) HeatingCoolingMode() model.HeatCoolMode {
    return model.HeatCoolMode(c.Byte())
}

// type CurrentRelativeHumidityCharacteristic struct {
//     *Float
//     humidity float64
// }
//
// func NewCurrentRelativeHumidityCharacteristic(value float64) *CurrentRelativeHumidityCharacteristic {
//     c := CurrentRelativeHumidityCharacteristic{NewFloat(value), value}
//     c.Type = CharTypeRelativeHumidityCurrent
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
//     c.Type = CharTypeRelativeHumidityTarget
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
//     c.Type = CharTypeCoolingThreshold
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
//     c.Type = CharTypeHeatingThreshold
//     c.Permissions = PermsAll()
//     return &c
// }