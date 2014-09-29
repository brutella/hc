package hk

type HeatingCoolingModeCharacteristic struct {
    *ByteCharacteristic
}

func NewHeatingCoolingModeCharacteristic(current HeatingCoolingMode, charType CharType, permissions []string) *HeatingCoolingModeCharacteristic {
    c := HeatingCoolingModeCharacteristic{NewByteCharacteristic(byte(current))}
    c.Type = charType
    c.Permissions = permissions
    return &c
}

func NewCurrentHeatingCoolingModeCharacteristic(current HeatingCoolingMode) *HeatingCoolingModeCharacteristic {
    return NewHeatingCoolingModeCharacteristic(current, CharTypeHeatingCoolingModeCurrent, PermsRead())
}

func NewTargetHeatingCoolingModeCharacteristic(current HeatingCoolingMode) *HeatingCoolingModeCharacteristic {
    return NewHeatingCoolingModeCharacteristic(current, CharTypeHeatingCoolingModeTarget, PermsRead())
}

func (c *HeatingCoolingModeCharacteristic) SetHeatingCoolingMode(mode HeatingCoolingMode) {
    c.SetByte(byte(mode))
}


func (c *HeatingCoolingModeCharacteristic) HeatingCoolingMode() HeatingCoolingMode {
    return HeatingCoolingMode(c.Byte())
}

// type CurrentRelativeHumidityCharacteristic struct {
//     *FloatCharacteristic
//     humidity float64
// }
//
// func NewCurrentRelativeHumidityCharacteristic(value float64) *CurrentRelativeHumidityCharacteristic {
//     c := CurrentRelativeHumidityCharacteristic{NewFloatCharacteristic(value), value}
//     c.Type = CharTypeRelativeHumidityCurrent
//     c.Permissions = PermsRead()
//     return &c
// }
//
// type TargetRelativeHumidityCharacteristic struct {
//     *FloatCharacteristic
//     target float64
// }
//
// func NewTargetRelativeHumidityCharacteristic(value, min, max, steps float64) *TargetRelativeHumidityCharacteristic {
//     c := TargetRelativeHumidityCharacteristic{NewFloatCharacteristicMinMaxSteps(value, min, max, steps), value}
//     c.Type = CharTypeRelativeHumidityTarget
//     c.Permissions = PermsAll()
//     return &c
// }
//
// type CoolingThresholdTemperatureCharacteristic struct {
//     *FloatCharacteristic
// }
//
// func NewCoolingThresholdTemperatureCharacteristic(value, min, max, steps float64) *CoolingThresholdTemperatureCharacteristic {
//     c := CoolingThresholdTemperatureCharacteristic{NewFloatCharacteristicMinMaxSteps(value, min, max, steps), value}
//     c.Type = CharTypeCoolingThreshold
//     c.Permissions = PermsAll()
//     return &c
// }
//
// type HeatingThresholdTemperatureCharacteristic struct {
//     *FloatCharacteristic
// }
//
// func NewHeatingThresholdTemperatureCharacteristic(value, min, max, steps float64) *HeatingThresholdTemperatureCharacteristic {
//     c := HeatingThresholdTemperatureCharacteristic{NewFloatCharacteristicMinMaxSteps(value, min, max, steps), value}
//     c.Type = CharTypeHeatingThreshold
//     c.Permissions = PermsAll()
//     return &c
// }