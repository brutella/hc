package accessory

import (
	"github.com/brutella/hc/service"
)

type GarageDoorOpener struct {
	*Accessory
	GarageDoorOpener *service.GarageDoorOpener
}

// NewGarageDoorOpener returns an accessory containing one service.
func NewGarageDoorOpener(info Info) *GarageDoorOpener {
	acc := GarageDoorOpener{}
	acc.Accessory = New(info, TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()

	acc.AddService(acc.GarageDoorOpener.Service)

	return &acc
}
