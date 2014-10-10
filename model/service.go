package model

type Service interface {
    Compareable
    
    SetId(int)
    GetCharacteristics()[]Characteristic
}