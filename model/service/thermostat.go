package service

import(
    "github.com/brutella/hap/model/characteristic"
    "github.com/brutella/hap/model"
)

type TempChangeFunc func(float64)
type Thermostat struct {
    *Service
    
    Name        *characteristic.Name
    Unit        *characteristic.TemperatureUnit
    Temp        *characteristic.TemperatureCharacteristic
    TargetTemp  *characteristic.TemperatureCharacteristic
    Mode        *characteristic.HeatingCoolingMode
    TargetMode  *characteristic.HeatingCoolingMode
    
    targetTempChange TempChangeFunc
}

// HomeKit does not support thermometers
// We use a thermostat with readonly services
// TODO File radar
func NewThermometer(name string, temperature float64) *Thermostat {
    thermostat := NewThermostat(name, temperature, 0, 100, 1)
    
    thermostat.TargetTemp.Permissions = characteristic.PermsRead()
    thermostat.TargetMode.Permissions = characteristic.PermsRead()
    
    return thermostat
}

func NewThermostat(name string, temperature, min, max, steps float64) *Thermostat {
    name_char  := characteristic.NewName(name)
    unit       := model.TempUnitCelsius
    unit_char  := characteristic.NewTemperatureUnit(unit)
    temp       := characteristic.NewCurrentTemperatureCharacteristic(temperature, min, max, steps, string(unit))
    targetTemp := characteristic.NewTargetTemperatureCharacteristic(temperature, min, max, steps, string(unit))
    mode       := characteristic.NewCurrentHeatingCoolingMode(model.ModeOff)
    targetMode := characteristic.NewTargetHeatingCoolingMode(model.ModeOff)
    
    service := NewService()
    service.Type = TypeThermostat
    service.AddCharacteristic(name_char.Characteristic)
    service.AddCharacteristic(unit_char.Characteristic)
    service.AddCharacteristic(temp.Characteristic)
    service.AddCharacteristic(targetTemp.Characteristic)
    service.AddCharacteristic(mode.Characteristic)
    service.AddCharacteristic(targetMode.Characteristic)
    
    t := Thermostat{service, name_char, unit_char, temp, targetTemp, mode, targetMode, nil}
    
    return &t
}