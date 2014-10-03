package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type Thermostat struct {
    *Service
    
    Name *characteristic.Name
    Unit *characteristic.TemperatureUnit
    Temp *characteristic.TemperatureCharacteristic
    TargetTemp *characteristic.TemperatureCharacteristic
    Mode *characteristic.HeatingCoolingMode
    TargetMode *characteristic.HeatingCoolingMode
}

func NewThermostat(name string, temperature, min, max, steps float64) *Thermostat {
    name_char  := characteristic.NewName(name)
    unit       := characteristic.UnitCelsius
    unit_char  := characteristic.NewTemperatureUnit(unit)
    temp       := characteristic.NewCurrentTemperatureCharacteristic(temperature, min, max, steps, unit)
    targetTemp := characteristic.NewTargetTemperatureCharacteristic(temperature, min, max, steps, unit)
    mode       := characteristic.NewCurrentHeatingCoolingMode(characteristic.ModeOff)
    targetMode := characteristic.NewTargetHeatingCoolingMode(characteristic.ModeOff)
    
    service := NewService()
    service.Type = TypeThermostat
    service.AddCharacteristic(name_char.Characteristic)
    service.AddCharacteristic(unit_char.Characteristic)
    service.AddCharacteristic(temp.Characteristic)
    service.AddCharacteristic(targetTemp.Characteristic)
    service.AddCharacteristic(mode.Characteristic)
    service.AddCharacteristic(targetMode.Characteristic)
    
    return &Thermostat{service, name_char, unit_char, temp, targetTemp, mode, targetMode}
}