package model

// A Compareable type can be compared.
type Compareable interface {
	// Equal returns true when argument is the same as the receiver
	Equal(interface{}) bool
}
