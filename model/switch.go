package model

type Switch interface {
    Accessory
    
    // Sets the switch state
    SetOn(on bool)
    
    // Returns the switch on state
    IsOn() bool
    
    // Sets the on state changed callback
    OnStateChanged(func(bool))
}