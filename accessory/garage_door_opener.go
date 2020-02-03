package accessory

import (
	"github.com/brutella/hc/service"
)

//GarageDoorOpener Struct
type GarageDoorOpener struct {
	*Accessory
	GarageDoorOpener *service.GarageDoorOpener
}

//NewGarageDoorOpener returns an garage accessory wh
func NewGarageDoorOpener(info Info) *GarageDoorOpener {
	acc := GarageDoorOpener{}
	acc.Accessory = New(info, TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()

	acc.AddService(acc.GarageDoorOpener.Service)

	return &acc
}
