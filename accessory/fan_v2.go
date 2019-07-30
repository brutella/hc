package accessory

import (
	"github.com/brutella/hc/service"
)

//FanV2 structure
type FanV2 struct {
	*Accessory
	FanV2 *service.FanV2
}

// NewFanV2 returns an outlet accessory containing one outlet service.
func NewFanV2(info Info) *FanV2 {
	acc := FanV2{}
	acc.Accessory = New(info, TypeFan)
	acc.FanV2 = service.NewFanV2()

	acc.AddService(acc.FanV2.Service)

	return &acc
}
