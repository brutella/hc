package model

type SwitchService struct {
    *Service
}

func NewSwitchService(name string, on bool) *SwitchService {
    on_char   := NewOnCharacteristic(on)
    name_char := NewNameCharacteristic(name)
    
    service := NewService()
    service.Type = ServiceTypeSwitch
    service.AddCharacteristic(on_char.Characteristic)
    service.AddCharacteristic(name_char.Characteristic)
    
    return &SwitchService{service}
}