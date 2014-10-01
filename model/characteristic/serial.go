package characteristic

type SerialNumber struct {
    *String
}

func NewSerialNumber(serial string) *SerialNumber {
    str := NewString(serial)
    str.Type = CharTypeSerialNumber
    str.Permissions = PermsReadOnly()
    
    return &SerialNumber{str}
}

func (c *SerialNumber) SetSerialNumber(serialNumber string) {
    c.SetString(serialNumber)
}

func (c *SerialNumber) SerialNumber() string {
    return c.StringValue()
}
