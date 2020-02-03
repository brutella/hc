package accessory

import (
	"github.com/brutella/hc/service"
)

//AirPurifier structure
type AirPurifier struct {
	*Accessory
	AirPurifier *service.AirPurifier
}

// NewAirPurifier returns an outlet accessory containing one outlet service.
func NewAirPurifier(info Info) *AirPurifier {
	acc := AirPurifier{}
	acc.Accessory = New(info, TypeAirPurifier)
	acc.AirPurifier = service.NewAirPurifier()

	acc.AddService(acc.AirPurifier.Service)

	return &acc
}
