package model

type Service interface {
    Compareable
    
    SetId(int64)
    GetId() int64
    GetCharacteristics()[]Characteristic
}