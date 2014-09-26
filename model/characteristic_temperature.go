package model


type TemperatureUnitCharacteristic struct {
    *ByteCharacteristic
}

func NewTemperatureUnitCharacteristic(unit string) *TemperatureUnitCharacteristic {
    b := ByteFromUnit(unit)
    c := TemperatureUnitCharacteristic{NewByteCharacteristic(b)}
    c.Type = CharTypeTemperatureUnits
    c.Permissions = PermsAll()
    return &c
}

type CurrentTemperatureCharacteristic struct {
    *FloatCharacteristic
    value float64
}

func NewCurrentTemperatureCharacteristic(value, min, max, steps float64, unit string) *CurrentTemperatureCharacteristic {
    c := CurrentTemperatureCharacteristic{NewFloatCharacteristicMinMaxSteps(value, min, max, steps), value}
    c.Unit = unit
    c.Type = CharTypeTemperatureCurrent
    c.Permissions = PermsRead()
    return &c
}

type TargetTemperatureCharacteristic struct {
    *FloatCharacteristic
    target float64
}

func NewTargetTemperatureCharacteristic(value, min, max, steps float64, unit string) *TargetTemperatureCharacteristic {
    c := TargetTemperatureCharacteristic{NewFloatCharacteristicMinMaxSteps(value, min, max, steps), value}
    c.Unit = unit
    c.Type = CharTypeTemperatureTarget
    c.Permissions = PermsAll()
    return &c
}

type ThermostatMode byte
const (
    ThermostatModeOff = 0x00
    ThermostatModeHeating = 0x01
    ThermostatModeCooling = 0x02
)

type TargetHeatingCoolingModeCharacteristic struct {
    *ByteCharacteristic
    target ThermostatMode
}

func NewTargetHeatingCoolingModeCharacteristic(target ThermostatMode) *TargetHeatingCoolingModeCharacteristic {
    c := TargetHeatingCoolingModeCharacteristic{NewByteCharacteristic(byte(target)), target}
    c.Type = CharTypeHeatingCoolingModeTarget
    c.Permissions = PermsAll()
    return &c
}

type CurrentHeatingCoolingModeCharacteristic struct {
    *ByteCharacteristic
    current ThermostatMode
}

func NewCurrentHeatingCoolingModeCharacteristic(current ThermostatMode) *CurrentHeatingCoolingModeCharacteristic {
    c := CurrentHeatingCoolingModeCharacteristic{NewByteCharacteristic(byte(current)), current}
    c.Type = CharTypeHeatingCoolingModeCurrent
    c.Permissions = PermsRead()
    return &c
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