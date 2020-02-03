package accessory

import (
	"github.com/brutella/hc/service"
)

//Lightswitch struct
type Lightswitch struct {
	*Accessory
	Lightswitch *service.Lightswitch
}

//NewLightswitch function
func NewLightswitch(info Info) *Lightswitch {
	acc := Lightswitch{}
	acc.Accessory = New(info, TypeLightbulb)
	acc.Lightswitch = service.NewLightswitch()

	acc.AddService(acc.Lightswitch.Service)

	return &acc
}
