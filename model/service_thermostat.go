package hk

type ThermostatService struct {
    *Service
    
    name *NameCharacteristic
    unit *TemperatureUnitCharacteristic
    temp *TemperatureCharacteristic
    targetTemp *TemperatureCharacteristic
    mode *HeatingCoolingModeCharacteristic
    targetMode *HeatingCoolingModeCharacteristic
}

/*
@property(retain, nonatomic) HAKNameCharacteristic *nameCharacteristic;
@property(readonly, nonatomic) HAKTemperatureUnitsCharacteristic *temperatureUnitsCharacteristic;
@property(readonly, nonatomic) HAKTargetTemperatureCharacteristic *targetTemperatureCharacteristic;
@property(readonly, nonatomic) HAKTemperatureCharacteristic *currentTemperatureCharacteristic;
@property(readonly, nonatomic) HAKTargetHeatingCoolingModeCharacteristic *targetHeatingCoolingModeCharacteristic;
@property(readonly, nonatomic) HAKHeatingCoolingModeCharacteristic *currentHeatingCoolingModeCharacteristic;
*/
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