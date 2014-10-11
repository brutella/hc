package accessory

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/service"
)

type switcher struct {
    *Accessory
    switcher *service.Switch
}

func NewSwitch(info model.Info) *switcher {
    accessory := New(info)
    s := service.NewSwitch(info.Name, false) // off
    
    accessory.AddService(s.Service)
    
    return &switcher{accessory, s}
}

func (s *switcher) SetOn(on bool) {
    s.switcher.On.SetOn(on)
}

func (s *switcher) IsOn() bool {
    return s.switcher.On.On()
}