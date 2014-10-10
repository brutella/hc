package model

type Switch interface {
    Accessory
    
    SetOn(on bool)
    GetOn() bool
}