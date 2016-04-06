package accessory

import (
	"github.com/brutella/hc/service"
)

type Switch struct {
	*Accessory
	Switch *service.Switch
}

// NewSwitch returns a switch which implements model.Switch.
func NewSwitch(info Info) *Switch {
	acc := Switch{}
	acc.Accessory = New(info, TypeOutlet)
	acc.Switch = service.NewSwitch()
	acc.AddService(acc.Switch.Service)

	return &acc
}
