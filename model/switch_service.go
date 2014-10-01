package model

type StateChangeFunc func(bool)
type SwitchService struct {
    *Service
    On *OnCharacteristic
    Name *NameCharacteristic
    fn StateChangeFunc
}

func NewSwitchService(name string, on bool) *SwitchService {
    on_char   := NewOnCharacteristic(on)
    name_char := NewNameCharacteristic(name)
    
    service := NewService()
    service.Type = ServiceTypeSwitch
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
func (s *SwitchService) CharactericDidChangeValue(c *Characteristic, change CharacteristicChange) {
    if s.fn != nil {
        s.fn(s.On.On())
    }
}