package accessory

import (
	"github.com/brutella/hc/service"
)

//LightDimer struct
type LightDimer struct {
	*Accessory
	LightDimer *service.LightDimer
}

//NewLightDimer function
func NewLightDimer(info Info) *LightDimer {
	acc := LightDimer{}
	acc.Accessory = New(info, TypeLightbulb)
	acc.LightDimer = service.NewLightDimer()

	acc.AddService(acc.LightDimer.Service)

	return &acc
}
