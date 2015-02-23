package model

// Compareables can be compared using the Equal method
type Compareable interface {
	// Equal returns true when argument is the same as the receiver
	Equal(interface{}) bool
}
