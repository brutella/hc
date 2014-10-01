package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type AccessoryInfoService struct {
    *Service
    
    Identify *characteristic.IdentifyCharacteristic
    Serial *characteristic.SerialNumberCharacteristic
    Model *characteristic.ModelCharacteristic
    Manufacturer *characteristic.ManufacturerCharacteristic
    Name *characteristic.NameCharacteristic
}

func NewAccessoryInfoService(accessoryName, serialNumber, manufacturerName, modelName string) *AccessoryInfoService {
    identify        := characteristic.NewIdentifyCharacteristic(false)
    serial          := characteristic.NewSerialNumberCharacteristic(serialNumber)
    model           := characteristic.NewModelCharacteristic(modelName)
    manufacturer    := characteristic.NewManufacturerCharacteristic(manufacturerName)
    name            := characteristic.NewNameCharacteristic(accessoryName)
    
    service := NewService()
    service.Type = TypeAccessoryInfo
    service.AddCharacteristic(identify.Characteristic)
    service.AddCharacteristic(serial.Characteristic)
    service.AddCharacteristic(model.Characteristic)
    service.AddCharacteristic(manufacturer.Characteristic)
    service.AddCharacteristic(name.Characteristic)
    
    return &AccessoryInfoService{service, identify, serial, model, manufacturer, name}
}