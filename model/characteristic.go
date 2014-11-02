package model

type Characteristic interface {
    Compareable
    
    // Returns the characteristic id
    GetId()int
    
    // Returns the raw value
    GetValue() interface{}
    
    // Sets the value
    // Only call this method when a client (e.g. iOS device) changes the value
    // Otherwise use the provided setter methods ( e.g. `switch.SetOn(true)`)
    SetValueFromRemote(interface{})
    
    // Enables/Disables events for this characteristic
    SetEventsEnabled(enable bool)
    
    // Returns true when events of this characteristic are enabled, otherwise false
    EventsEnabled() bool
}