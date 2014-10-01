package model

type ThermostatService struct {
    *Service
    
    Name *NameCharacteristic
    Unit *TemperatureUnitCharacteristic
    Temp *TemperatureCharacteristic
    TargetTemp *TemperatureCharacteristic
    Mode *HeatingCoolingModeCharacteristic
    TargetMode *HeatingCoolingModeCharacteristic
}

func NewThermostatService(name string, temperature, min, max, steps float64) *ThermostatService {
    name_char  := NewNameCharacteristic(name)
    unit       := UnitCelsius
    unit_char  := NewTemperatureUnitCharacteristic(unit)
    temp       := NewCurrentTemperatureCharacteristic(temperature, min, max, steps, unit)
    targetTemp := NewTargetTemperatureCharacteristic(temperature, min, max, steps, unit)
    mode       := NewCurrentHeatingCoolingModeCharacteristic(ModeOff)
    targetMode := NewTargetHeatingCoolingModeCharacteristic(ModeOff)
    
    service := NewService()
    service.Type = SerivceTypeThermostat
    service.AddCharacteristic(name_char.Characteristic)
    service.AddCharacteristic(unit_char.Characteristic)
    service.AddCharacteristic(temp.Characteristic)
    service.AddCharacteristic(targetTemp.Characteristic)
    service.AddCharacteristic(mode.Characteristic)
    service.AddCharacteristic(targetMode.Characteristic)
    
    return &ThermostatService{service, name_char, unit_char, temp, targetTemp, mode, targetMode}
}