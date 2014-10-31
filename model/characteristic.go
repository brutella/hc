package model

type Characteristic interface {
    Compareable
    
    // Returns the characteristic id
    GetId()int
    
    // Returns the value of the characteristic
    GetValue() interface{}
    
    // Sets the value
    // Only call this method when a client (e.g. iOS device) invokes
    // a value change. Otherwise use the accessory setter methods ( e.g. `switch.SetOn(true)`)
    SetValueFromRemote(interface{})
    
    // Enables or disables events for this characteristic
    EnableEvents(enable bool)
}