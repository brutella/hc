package model

type ThermostatService struct {
    *Service
}

/*
@property(retain, nonatomic) HAKCoolingThresholdTemperatureCharacteristic
@property(retain, nonatomic) HAKHeatingThresholdTemperatureCharacteristic 
@property(retain, nonatomic) HAKNameCharacteristic *nameCharacteristic;
@property(retain, nonatomic) HAKTargetRelativeHumidityCharacteristic *targetRelativeHumidityCharacteristic;
@property(retain, nonatomic) HAKCurrentRelativeHumidityCharacteristic *currentRelativeHumidityCharacteristic;
@property(readonly, nonatomic) HAKTemperatureUnitsCharacteristic *temperatureUnitsCharacteristic;
@property(readonly, nonatomic) HAKTargetTemperatureCharacteristic *targetTemperatureCharacteristic;
@property(readonly, nonatomic) HAKCurrentTemperatureCharacteristic *currentTemperatureCharacteristic;
@property(readonly, nonatomic) HAKTargetHeatingCoolingModeCharacteristic *targetHeatingCoolingModeCharacteristic;
@property(readonly, nonatomic) HAKCurrentHeatingCoolingModeCharacteristic *currentHeatingCoolingModeCharacteristic;
*/
func NewThermostatService(name string, temperature, min, max, steps float64) *ThermostatService {
    char_name := NewNameCharacteristic(name)
    unit := NewTemperatureUnitCharacteristic(UnitCelsius)
    current := NewCurrentTemperatureCharacteristic(temperature, min, max, steps)
    
    service := NewService()
    service.Type = SerivceTypeThermostat
    service.AddCharacteristic(char_name.Characteristic)
    service.AddCharacteristic(unit.Characteristic)
    service.AddCharacteristic(current.Characteristic)
    
    return &ThermostatService{service}
}