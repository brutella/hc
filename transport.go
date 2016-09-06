package hc

import (
	"github.com/brutella/hc/accessory"
)

// Transport provides accessories over a network.
type Transport interface {
	// Start starts the transport
	Start()

	// Stop stops the transport
	Stop()
	
	// Add an accessory to the transport - must be done before the transport is started
	AddAccessory(a *accessory.Accessory) error
	
	// Find an accessory by it's ID
	GetAccessoryByID(aid int64) (a *accessory.Accessory)
	
	// Get the container holding all the accessories
	GetContainer() (container *accessory.Container)
	
	// Delete data related to this transport from disk
	RemoveFromDisk() error
}
