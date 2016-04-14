// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSerialNumber = "30"

type SerialNumber struct {
	*String
}

func NewSerialNumber() *SerialNumber {
	char := NewString(TypeSerialNumber)
	char.Format = FormatString
	char.Perms = []string{PermRead}

	char.SetValue("")

	return &SerialNumber{char}
}
