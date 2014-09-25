package model

type ManufacturerCharacteristic struct {
    *StringCharacteristic
}

func NewManufacturerCharacteristic(manufacturerName string) *ManufacturerCharacteristic {
    str := NewStringCharacteristic(manufacturerName)
    str.Type = CharTypeManufacturer
    str.Permissions = []string{PermRead}
    
    return &ManufacturerCharacteristic{str}
}