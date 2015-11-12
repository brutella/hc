package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type switcher struct {
	*Accessory
	switcher *service.Switch
}

// NewSwitch returns a switch which implements model.Switch.
func NewSwitch(info model.Info) *switcher {
	accessory := New(info)
	s := service.NewSwitch(info.Name, false) // off

	accessory.AddService(s.Service)

	sw := switcher{accessory, s}

	return &sw
}

func (s *switcher) SetOn(on bool) {
	s.switcher.On.SetOn(on)
}

func (s *switcher) IsOn() bool {
	return s.switcher.On.On()
}

func (s *switcher) OnStateChanged(fn func(bool)) {
	s.switcher.On.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(bool))
	})
}
