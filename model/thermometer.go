package model

type TempUnit string
const(
    TempUnitCelsius = TempUnit("celsius")
)

type Thermometer interface {
    Accessory
    
    SetTemperature(float64)
    Temperature() float64
    Unit() TempUnit
    SetTargetTemperature(float64)
    TargetTemperature() float64
}