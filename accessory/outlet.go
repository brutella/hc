package accessory

import (
	"github.com/brutella/hc/service"
)

type Outlet struct {
	*Accessory
	Outlet *service.Outlet
}

// NewOutlet returns an outlet accessory containing one outlet service.
func NewOutlet(info Info) *Outlet {
	acc := Outlet{}
	acc.Accessory = New(info, TypeOutlet)
	acc.Outlet = service.NewOutlet()
	acc.Outlet.OutletInUse.SetValue(true)

	acc.AddService(acc.Outlet.Service)

	return &acc
}
