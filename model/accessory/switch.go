package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type Switch struct {
	*Accessory
	Switch *service.Switch
}

// NewSwitch returns a switch which implements model.Switch.
func NewSwitch(info model.Info) *Switch {
	accessory := New(info)
	s := service.NewSwitch(info.Name, false) // off

	accessory.AddService(s.Service)

	sw := Switch{accessory, s}

	return &sw
}

func (s *Switch) SetOn(on bool) {
	s.Switch.On.SetOn(on)
}

func (s *Switch) IsOn() bool {
	return s.Switch.On.On()
}

func (s *Switch) OnStateChanged(fn func(bool)) {
	s.Switch.On.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(bool))
	})
}
