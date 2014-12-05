package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type Switch struct {
    *Service
    On   *characteristic.On
    Name *characteristic.Name
}

func NewSwitch(name string, on bool) *Switch {
    on_char   := characteristic.NewOn(on)
    name_char := characteristic.NewName(name)
    
    service := New()
    service.Type = TypeSwitch
    service.AddCharacteristic(on_char.Characteristic)
    service.AddCharacteristic(name_char.Characteristic)
    
    return &Switch{service, on_char, name_char}
}