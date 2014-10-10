package service

import(
    "github.com/brutella/hap/model/characteristic"
)
type TempChangeFunc func(float64)
type Thermostat struct {
    *Service
    
    Name *characteristic.Name
    Unit *characteristic.TemperatureUnit
    Temp *characteristic.TemperatureCharacteristic
    TargetTemp *characteristic.TemperatureCharacteristic
    Mode *characteristic.HeatingCoolingMode
    TargetMode *characteristic.HeatingCoolingMode
    
    targetTempChange TempChangeFunc
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
    
    t := Thermostat{service, name_char, unit_char, temp, targetTemp, mode, targetMode, nil}
    
    targetTemp.AddRemoteChangeDelegate(&t)
    
    return &t
}

func (t *Thermostat) SetTemperature(value float64){
    t.Temp.SetTemperature(value)
}


func (t *Thermostat) TargetTempChanged(fn TempChangeFunc){
    t.targetTempChange = fn
}

func (t *Thermostat) CharactericDidChangeValue(c *characteristic.Characteristic, change characteristic.CharacteristicChange) {
    if c.Equal(t.TargetTemp) {
        if t.targetTempChange != nil {
            t.targetTempChange(t.TargetTemp.Temperature())
        }
    }
}