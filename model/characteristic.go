package model

// A Characteristic is identifiable and has a (observeable) value.
type Characteristic interface {
	Compareable

	// GetID returns the characteristic id
	GetID() int64

	// GetValue returns the characteristic value
	GetValue() interface{}

	// SetValueFromRemote sets the characteristic value
	// Only call this method when a client changes the value
	// Otherwise use the provided setter methods ( e.g. `switch.SetOn(true)`)
	SetValueFromRemote(interface{})

	// SetEventsEnabled dis-/enables events for the characteristic
	SetEventsEnabled(enable bool)

	// EventsEnabled returns if characteristic has event notifications enabled
	EventsEnabled() bool
}
