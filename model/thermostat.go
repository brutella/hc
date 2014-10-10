package model

type HeatCoolMode byte

const (
    // TODO verify the values
    ModeOff = HeatCoolMode(0x00)
    ModeHeating = HeatCoolMode(0x01)
    ModeCooling = HeatCoolMode(0x02)
)

type Thermostat interface {
    Thermometer
    
    Mode() HeatCoolMode
    SetTargetMode(HeatCoolMode)
    TargetMode() HeatCoolMode
}

type Hygrometer interface {
    Humidity() float64
    
    SetTargetHumidity(float64)
    TargetHumidity() float64
}