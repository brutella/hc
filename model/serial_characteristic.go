package hk

type SerialNumberCharacteristic struct {
    *StringCharacteristic
}

func NewSerialNumberCharacteristic(serial string) *SerialNumberCharacteristic {
    str := NewStringCharacteristic(serial)
    str.Type = CharTypeSerialNumber
    str.Permissions = PermsReadOnly()
    
    return &SerialNumberCharacteristic{str}
}

func (c *SerialNumberCharacteristic) SetSerialNumber(serialNumber string) {
    c.SetString(serialNumber)
}

func (c *SerialNumberCharacteristic) SerialNumber() string {
    return c.String()
}
