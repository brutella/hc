package model

type HeatCoolMode byte

const (
    // TODO verify the values
    ModeOff = HeatCoolMode(0x00)
    ModeHeating = HeatCoolMode(0x01)
    ModeCooling = HeatCoolMode(0x02)
)

// A thermostat measures temperature and changes the 
// mode (heating or cooling) to reach certain target temperature
type Thermostat interface {
    Thermometer
    
    SetTargetTemperature(float64)
    TargetTemperature() float64
    Mode() HeatCoolMode
    SetTargetMode(HeatCoolMode)
    TargetMode() HeatCoolMode
}