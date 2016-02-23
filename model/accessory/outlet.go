package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type Outlet struct {
	*Accessory
	Outlet *service.Outlet
}

// NewOutlet returns an Outlet which implements model.Outlet.
func NewOutlet(info model.Info) *Outlet {
	accessory := New(info)
	s := service.NewOutlet(info.Name, false, true) // off

	accessory.AddService(s.Service)

	sw := Outlet{accessory, s}

	return &sw
}

func (o *Outlet) SetOn(on bool) {
	o.Outlet.On.SetOn(on)
}

func (o *Outlet) IsOn() bool {
	return o.Outlet.On.On()
}

func (o *Outlet) SetInUse(on bool) {
	o.Outlet.InUse.SetInUse(on)
}

func (o *Outlet) IsInUse() bool {
	return o.Outlet.InUse.InUse()
}

func (o *Outlet) OnStateChanged(fn func(bool)) {
	o.Outlet.On.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(bool))
	})
}

func (o *Outlet) InUseStateChanged(fn func(bool)) {
	o.Outlet.InUse.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(bool))
	})
}
