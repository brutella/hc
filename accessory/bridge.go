package accessory

import (
	"github.com/brutella/hc/service"
)

type Bridge struct {
	*Accessory
	BridgingState *service.BridgingState
}

// NewBridge returns a bridge which implements model.Bridge.
func NewBridge(info Info, identifier string) *Bridge {
	acc := Bridge{}
	acc.Accessory = New(info, TypeBridge)
	acc.BridgingState = service.NewBridgingState()
	acc.BridgingState.Category.SetValue(1)
	acc.BridgingState.AccessoryIdentifier.SetValue(identifier)
	acc.BridgingState.Reachable.SetValue(true)
	acc.BridgingState.LinkQuality.SetValue(100)
	
	acc.AddService(acc.BridgingState.Service)

	return &acc
}
