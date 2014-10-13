package model

// Base interface for all accessories
type Accessory interface {
    Compareable
    
    // Returns the accessories id
    GetId()int
    
    // Returns the services which represent the accessory 
    GetServices()[]Service
    
    // Returns the name of the accessory
    Name() string
    
    // Returns the serial number of the accessory
    SerialNumber() string
    
    // Returns the manufacturer name of the accessory
    Manufacturer() string
    
    // Returns the model description of the accessory
    Model() string
}