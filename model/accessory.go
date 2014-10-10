package model

type Accessory interface {
    Compareable
    
    SetId(int)
    GetId()int
    
    GetServices()[]Service
    
    GetName() string
    GetSerialNumber() string
    GetManufacturer() string
    GetModel() string
}