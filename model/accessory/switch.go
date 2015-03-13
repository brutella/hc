package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
)

type switcher struct {
	*Accessory
	switcher *service.Switch

	onChanged func(bool)
}

// NewSwitch returns a switch which implements model.Switch.
func NewSwitch(info model.Info) *switcher {
	accessory := New(info)
	s := service.NewSwitch(info.Name, false) // off

	accessory.AddService(s.Service)

	sw := switcher{accessory, s, nil}

	s.On.OnRemoteChange(func(c *characteristic.Characteristic, new, old interface{}) {
		if sw.onChanged != nil {
			sw.onChanged(s.On.On())
		}
	})

	return &sw
}

func (s *switcher) SetOn(on bool) {
	s.switcher.On.SetOn(on)
}

func (s *switcher) IsOn() bool {
	return s.switcher.On.On()
}

func (s *switcher) OnStateChanged(fn func(bool)) {
	s.onChanged = fn
}
