package model

type Accessory interface {
    Compareable
    
    SetId(int)
    GetId()int
    
    GetServices()[]Service
    
    Name() string
    SerialNumber() string
    Manufacturer() string
    Model() string
}