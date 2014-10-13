package model

type Switch interface {
    Accessory
    
    // Changes the switche state to *on* or *off*
    SetOn(on bool)
    
    // Returns `true` when the switch is set to *on*, otherwise `false`
    IsOn() bool
    
    // Adds a function which is called when a client changed the on state
    // The function is not invoked when calling `SetOn`.
    OnStateChanged(func(bool))
}