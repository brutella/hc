package accessory

import (
	"github.com/brutella/hc/service"
)

type ColoredLightbulb struct {
	*Accessory
	Lightbulb *service.ColoredLightbulb
}

// NewLightbulb returns an light bulb accessory which one light bulb service.
func NewColoredLightbulb(info Info) *ColoredLightbulb {
	acc := ColoredLightbulb{}
	acc.Accessory = New(info, TypeLightbulb)
	acc.Lightbulb = service.NewColoredLightbulb()

	acc.Lightbulb.Brightness.SetValue(100)

	acc.AddService(acc.Lightbulb.Service)

	return &acc
}
