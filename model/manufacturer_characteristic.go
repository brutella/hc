package hk

type ManufacturerCharacteristic struct {
    *StringCharacteristic
}

func NewManufacturerCharacteristic(name string) *ManufacturerCharacteristic {
    str := NewStringCharacteristic(name)
    str.Type = CharTypeManufacturer
    str.Permissions = PermsReadOnly()
    
    return &ManufacturerCharacteristic{str}
}

func (c *ManufacturerCharacteristic) SetManufacturer(name string) {
    c.SetString(name)
}

func (c *ManufacturerCharacteristic) Manufacturer() string {
    return c.String()
}