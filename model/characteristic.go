package model

type Characteristic interface {
	Compareable

	// GetId returns the characteristic's id
	GetId() int64

	// GetValue returns the raw value
	GetValue() interface{}

	// SetValueFromRemote sets the value
	// Only call this method when a client (e.g. iOS device) changes the value
	// Otherwise use the provided setter methods ( e.g. `switch.SetOn(true)`)
	SetValueFromRemote(interface{})

	// SetEventsEnabled dis-/enables events for this characteristic
	SetEventsEnabled(enable bool)

	// EventsEnabled returns true when events of this characteristic are enabled, otherwise false
	EventsEnabled() bool
}
