package accessory

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/service"
    "github.com/brutella/hap/model/characteristic"
)

type switcher struct {
    *Accessory
    switcher *service.Switch
    
    onChanged func(bool)
}

func NewSwitch(info model.Info) *switcher {
    accessory := New(info)
    s := service.NewSwitch(info.Name, false) // off
    
    accessory.AddService(s.Service)
    
    sw := switcher{accessory, s, nil}
    s.On.AddRemoteChangeDelegate(&sw)
        
    return &sw
}

func (s *switcher) SetOn(on bool) {
    s.switcher.On.SetOn(on)
}

func (s *switcher) IsOn() bool {
    return s.switcher.On.On()
}

func (s *switcher) OnStateChanged(fn func(bool)){
    s.onChanged = fn
}

func (s *switcher) CharactericDidChangeValue(c *characteristic.Characteristic, change characteristic.CharacteristicChange) {
    if s.onChanged != nil {
        s.onChanged(s.switcher.On.On())
    }
}