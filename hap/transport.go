package hap

import ()

// Transport provides ressources over a network
type Transport interface {
	// Start starts the transpor.t
	Start()

	// Stop stops the transport.
	Stop()
}
