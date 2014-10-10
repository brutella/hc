package model

type Characteristic interface {
    Compareable
    
    SetId(int)
    GetId()int
    
    GetValue() interface{}
    SetValueFromRemote(interface{})
}