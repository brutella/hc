package accessory

import (
	"github.com/brutella/hc/service"
)

type HeaterCooler struct {
	*Accessory
	HeaterCooler *service.HeaterCooler
}

// NewSwitch returns a heater which implements model.Heater.
func NewHeater(info Info) *HeaterCooler {
	acc := HeaterCooler{}
	acc.Accessory = New(info, TypeHeater)
	acc.HeaterCooler = service.NewHeaterCooler()
	acc.AddService(acc.HeaterCooler.Service)

	return &acc
}
