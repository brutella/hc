package model

type ThermostatService struct {
    *Service
}

/*
@property(retain, nonatomic) HAKNameCharacteristic *nameCharacteristic;
@property(readonly, nonatomic) HAKTemperatureUnitsCharacteristic *temperatureUnitsCharacteristic;
@property(readonly, nonatomic) HAKTargetTemperatureCharacteristic *targetTemperatureCharacteristic;
@property(readonly, nonatomic) HAKCurrentTemperatureCharacteristic *currentTemperatureCharacteristic;
@property(readonly, nonatomic) HAKTargetHeatingCoolingModeCharacteristic *targetHeatingCoolingModeCharacteristic;
@property(readonly, nonatomic) HAKCurrentHeatingCoolingModeCharacteristic *currentHeatingCoolingModeCharacteristic;
*/
func NewThermostatService(name string, temperature, min, max, steps float64) *ThermostatService {
    char_name   := NewNameCharacteristic(name)
    unit := UnitCelsius
    unit_char    := NewTemperatureUnitCharacteristic(unit)
    current      := NewCurrentTemperatureCharacteristic(temperature, min, max, steps, unit)
    target      := NewTargetTemperatureCharacteristic(temperature, min, max, steps, unit)
    current_mode := NewCurrentHeatingCoolingModeCharacteristic(ThermostatModeOff)
    target_mode := NewTargetHeatingCoolingModeCharacteristic(ThermostatModeOff)
    
    
    service := NewService()
    service.Type = SerivceTypeThermostat
    service.AddCharacteristic(char_name.Characteristic)
    service.AddCharacteristic(unit_char.Characteristic)
    service.AddCharacteristic(current.Characteristic)
    service.AddCharacteristic(target.Characteristic)
    service.AddCharacteristic(current_mode.Characteristic)
    service.AddCharacteristic(target_mode.Characteristic)
    
    return &ThermostatService{service}
}