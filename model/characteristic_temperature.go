package model

type TemperatureUnitCharacteristic struct {
    *StringCharacteristic
}

func NewTemperatureUnitCharacteristic(unit string) *TemperatureUnitCharacteristic {
    return &TemperatureUnitCharacteristic{NewStringCharacteristic(unit)}
}

type CurrentTemperatureCharacteristic struct {
    *FloatCharacteristic
    value float64
}

func NewCurrentTemperatureCharacteristic(value, min, max, steps float64) *CurrentTemperatureCharacteristic {
    return &CurrentTemperatureCharacteristic{NewFloatCharacteristic(value, min, max, steps), value}
}
