package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type AccessoryInfo struct {
    *Service
    
    Identify *characteristic.Identify
    Serial *characteristic.SerialNumber
    Model *characteristic.Model
    Manufacturer *characteristic.Manufacturer
    Name *characteristic.Name
}

func NewAccessoryInfo(accessoryName, serialNumber, manufacturerName, modelName string) *AccessoryInfo {
    identify        := characteristic.NewIdentify(false)
    serial          := characteristic.NewSerialNumber(serialNumber)
    model           := characteristic.NewModel(modelName)
    manufacturer    := characteristic.NewManufacturer(manufacturerName)
    name            := characteristic.NewName(accessoryName)
    
    service := NewService()
    service.Type = TypeAccessoryInfo
    service.AddCharacteristic(identify.Characteristic)
    service.AddCharacteristic(serial.Characteristic)
    service.AddCharacteristic(model.Characteristic)
    service.AddCharacteristic(manufacturer.Characteristic)
    service.AddCharacteristic(name.Characteristic)
    
    return &AccessoryInfo{service, identify, serial, model, manufacturer, name}
}