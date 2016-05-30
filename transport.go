package hc

// Transport provides accessories over a network.
type Transport interface {
	// Start starts the transport
	Start()

	// Stop stops the transport
	Stop()
}
