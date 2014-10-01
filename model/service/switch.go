package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type StateChangeFunc func(bool)
type Switch struct {
    *Service
    On *characteristic.On
    Name *characteristic.Name
    fn StateChangeFunc
}

func NewSwitch(name string, on bool) *Switch {
    on_char   := characteristic.NewOn(on)
    name_char := characteristic.NewName(name)
    
    service := NewService()
    service.Type = TypeSwitch
    service.AddCharacteristic(on_char.Characteristic)
    service.AddCharacteristic(name_char.Characteristic)
    
    s := &Switch{service, on_char, name_char, nil}
    
    on_char.AddRemoteChangeDelegate(s)
    
    return s
}

func (s *Switch) OnStateChanged(fn StateChangeFunc){
    s.fn = fn
}

// On Characteristic Changed
func (s *Switch) CharactericDidChangeValue(c *characteristic.Characteristic, change characteristic.CharacteristicChange) {
    if s.fn != nil {
        s.fn(s.On.On())
    }
}