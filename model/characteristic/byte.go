package characteristic

type ByteCharacteristic struct {
	*Number
}

func NewByteCharacteristic(value byte, permissions []string) *ByteCharacteristic {
	number := NewNumber(value, nil, nil, nil, FormatByte, permissions)
	return &ByteCharacteristic{number}
}

func (c *ByteCharacteristic) SetByte(value byte) {
	c.SetNumber(value)
}

func (c *ByteCharacteristic) Byte() byte {
	return c.GetValue().(byte)
}
