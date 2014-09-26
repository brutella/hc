package model

type ManufacturerCharacteristic struct {
    *StringCharacteristic
}

func NewManufacturerCharacteristic(manufacturerName string) *ManufacturerCharacteristic {
    str := NewStringCharacteristic(manufacturerName)
    str.Type = CharTypeManufacturer
    str.Permissions = PermsReadOnly()
    
    return &ManufacturerCharacteristic{str}
}