package service

import(
    "github.com/brutella/hap/model/characteristic"
)

type StateChangeFunc func(bool)
type SwitchService struct {
    *Service
    On *characteristic.OnCharacteristic
    Name *characteristic.NameCharacteristic
    fn StateChangeFunc
}

func NewSwitchService(name string, on bool) *SwitchService {
    on_char   := characteristic.NewOnCharacteristic(on)
    name_char := characteristic.NewNameCharacteristic(name)
    
    service := NewService()
    service.Type = TypeSwitch
    service.AddCharacteristic(on_char.Characteristic)
    service.AddCharacteristic(name_char.Characteristic)
    
    s := &SwitchService{service, on_char, name_char, nil}
    
    on_char.AddRemoteChangeDelegate(s)
    
    return s
}

func (s *SwitchService) OnStateChanged(fn StateChangeFunc){
    s.fn = fn
}

// On Characteristic Changed
func (s *SwitchService) CharactericDidChangeValue(c *characteristic.Characteristic, change characteristic.CharacteristicChange) {
    if s.fn != nil {
        s.fn(s.On.On())
    }
}