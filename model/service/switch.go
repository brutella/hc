package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type OnChangeFunc func(bool)
type Switch struct {
    *Service
    On *characteristic.On
    Name *characteristic.Name
    
    fn OnChangeFunc
}

func NewSwitch(name string, on bool) *Switch {
    on_char   := characteristic.NewOn(on)
    name_char := characteristic.NewName(name)
    
    service := NewService()
    service.Type = TypeSwitch
    service.AddCharacteristic(on_char)
    service.AddCharacteristic(name_char)
    
    s := Switch{service, on_char, name_char, nil}
    
    on_char.AddRemoteChangeDelegate(&s)
    
    return &s
}

func (s *Switch) OnStateChanged(fn OnChangeFunc){
    s.fn = fn
}

// On Characteristic Changed
func (s *Switch) CharactericDidChangeValue(c *characteristic.Characteristic, change characteristic.CharacteristicChange) {
    if s.fn != nil {
        s.fn(s.On.On())
    }
}