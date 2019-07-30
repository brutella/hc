package accessory

import (
	"github.com/brutella/hc/service"
)

//Fan struct
type Fan struct {
	*Accessory
	Fan *service.Fan
}

//NewFan function
func NewFan(info Info) *Fan {
	acc := Fan{}
	acc.Accessory = New(info, TypeFan)
	acc.Fan = service.NewFan()

	acc.AddService(acc.Fan.Service)

	return &acc
}
