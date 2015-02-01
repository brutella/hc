package characteristic

type String struct {
	*Characteristic
}

func NewString(value string) *String {
	return &String{NewCharacteristic(value, FormatString, CharTypeUnknown, nil)}
}

func (s *String) SetString(str string) {
	s.SetValue(str)
}

func (s *String) StringValue() string {
	return s.Value.(string)
}
