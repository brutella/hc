package model

type HeatCoolMode byte

const (
    // TODO verify the values
    ModeOff = HeatCoolMode(0x00)
    ModeHeating = HeatCoolMode(0x01)
    ModeCooling = HeatCoolMode(0x02)
    ModeAuto = HeatCoolMode(0x03)
)

// A thermostat measures temperature and changes the 
// mode (heating, cooling, auto) to reach certain target temperature
type Thermostat interface {
    Thermometer
    
    // Sets the target temperature
    SetTargetTemperature(float64)
    
    // Returns the target temperature
    TargetTemperature() float64
    
    // SetMode sets the current mode
    // ModeAuto is ignored because the current mode cannot be auto
    SetMode(HeatCoolMode)
    
    // Returns the mode
    Mode() HeatCoolMode
    
    // Sets the target mode
    SetTargetMode(HeatCoolMode)
    
    // Returns the target mode
    TargetMode() HeatCoolMode
}