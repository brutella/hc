package hk

type AccessoryInfoService struct {
    *Service
}

func NewAccessoryInfoService(accessoryName, serialNumber, manufacturerName, modelName string) *AccessoryInfoService {
    identify        := NewIdentifyCharacteristic(false)
    serial          := NewSerialNumberCharacteristic(serialNumber)
    model           := NewModelCharacteristic(modelName)
    manufacturer    := NewManufacturerCharacteristic(manufacturerName)
    name            := NewNameCharacteristic(accessoryName)
    
    service := NewService()
    service.Type = ServiceTypeAccessoryInfo
    service.AddCharacteristic(identify.Characteristic)
    service.AddCharacteristic(serial.Characteristic)
    service.AddCharacteristic(model.Characteristic)
    service.AddCharacteristic(manufacturer.Characteristic)
    service.AddCharacteristic(name.Characteristic)
    
    return &AccessoryInfoService{service}
}