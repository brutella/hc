package model

// A Switch is an accessory which has an on state
type Switch interface {
	Accessory

	// SetOn sets the switch state to the argument boolean
	SetOn(on bool)

	// IsOn return true when the switch is on, otherwise false
	IsOn() bool

	// OnStateChanged calls the argument function when the switch's
	// on state changed
	OnStateChanged(func(bool))
}
