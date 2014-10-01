package characteristic

type ByteCharacteristic struct {
    *NumberCharacteristic
}

func NewByteCharacteristic(value byte) *ByteCharacteristic {
    number := NewNumberCharacteristic(value, nil, nil, nil, FormatByte)
    return &ByteCharacteristic{number}
}

func (c *ByteCharacteristic) SetByte(value byte) {
    c.SetValue(value)
}

func (c *ByteCharacteristic) Byte() byte {
    return c.GetValue().(byte)
}