package model

type AccessoryInfoService struct {
    *Service
}

func NewAccessoryInfoService(serialNumber, modelName, manufacturerName, accessoryName string) *AccessoryInfoService {
    identify        := NewIdentifyCharacteristic(false)
    serial          := NewSerialNumberCharacteristic(serialNumber)
    model           := NewModelCharacteristic(modelName)
    manufacturer    := NewManufacturerCharacteristic(manufacturerName)
    name            := NewNameCharacteristic(accessoryName)
    // characteristics := []interface{}{identify, serial, model, manufacturer, name}
    service := NewService()
    service.Type = ServiceTypeAccessoryInfo
    service.AddCharacteristic(identify.Characteristic)
    service.AddCharacteristic(serial.Characteristic)
    service.AddCharacteristic(model.Characteristic)
    service.AddCharacteristic(manufacturer.Characteristic)
    service.AddCharacteristic(name.Characteristic)
    
    return &AccessoryInfoService{service}
}