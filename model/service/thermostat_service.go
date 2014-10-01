package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type ThermostatService struct {
    *Service
    
    Name *characteristic.NameCharacteristic
    Unit *characteristic.TemperatureUnitCharacteristic
    Temp *characteristic.TemperatureCharacteristic
    TargetTemp *characteristic.TemperatureCharacteristic
    Mode *characteristic.HeatingCoolingModeCharacteristic
    TargetMode *characteristic.HeatingCoolingModeCharacteristic
}

func NewThermostatService(name string, temperature, min, max, steps float64) *ThermostatService {
    name_char  := characteristic.NewNameCharacteristic(name)
    unit       := characteristic.UnitCelsius
    unit_char  := characteristic.NewTemperatureUnitCharacteristic(unit)
    temp       := characteristic.NewCurrentTemperatureCharacteristic(temperature, min, max, steps, unit)
    targetTemp := characteristic.NewTargetTemperatureCharacteristic(temperature, min, max, steps, unit)
    mode       := characteristic.NewCurrentHeatingCoolingModeCharacteristic(characteristic.ModeOff)
    targetMode := characteristic.NewTargetHeatingCoolingModeCharacteristic(characteristic.ModeOff)
    
    service := NewService()
    service.Type = TypeThermostat
    service.AddCharacteristic(name_char.Characteristic)
    service.AddCharacteristic(unit_char.Characteristic)
    service.AddCharacteristic(temp.Characteristic)
    service.AddCharacteristic(targetTemp.Characteristic)
    service.AddCharacteristic(mode.Characteristic)
    service.AddCharacteristic(targetMode.Characteristic)
    
    return &ThermostatService{service, name_char, unit_char, temp, targetTemp, mode, targetMode}
}