package model

type Switch interface {
    Accessory
    
    SetOn(on bool)
    IsOn() bool
    OnStateChanged(func(bool))
}