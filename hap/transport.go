package hap

import ()

// OnStopFunc is the function which is invoked when the transport stops.
type OnStopFunc func()

// Transport provides ressources over a network
type Transport interface {
	// Start starts the transpor.t
	Start()

	// Stop stops the transport.
	Stop()

	// OnStop calls a function when the transport is stopped.
	OnStop(fn OnStopFunc)
}
