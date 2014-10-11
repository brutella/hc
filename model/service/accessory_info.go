package service

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/characteristic"
)

type AccessoryInfo struct {
    *Service
    
    Identify *characteristic.Identify           `json:"-"`
    Serial *characteristic.SerialNumber         `json:"-"`
    Model *characteristic.Model                 `json:"-"`
    Manufacturer *characteristic.Manufacturer   `json:"-"`
    Name *characteristic.Name                   `json:"-"`
}

func NewInfo(info model.Info) *AccessoryInfo {
    return NewAccessoryInfo(info.Name, info.Serial, info.Manufacturer, info.Model)
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