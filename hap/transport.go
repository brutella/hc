package hap

import ()

// OnStopFunc is the function which is invoked when the transport stops.
type OnStopFunc func()

type Transport interface {
	Start()
	Stop()
	OnStop(fn OnStopFunc)
}
