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

	// Add accessory to transport
	addAccessory(a *accessory.Accessory)
}
