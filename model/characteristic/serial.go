package characteristic

type SerialNumber struct {
	*String
}

func NewSerialNumber(serial string) *SerialNumber {
	str := NewString(serial)
	str.Type = TypeSerialNumber
	str.Permissions = PermsReadOnly()

	return &SerialNumber{str}
}

func (s *SerialNumber) SetSerialNumber(serialNumber string) {
	s.SetString(serialNumber)
}

func (s SerialNumber) SerialNumber() string {
	return s.StringValue()
}
