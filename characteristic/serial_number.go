// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSerialNumber = "00000030-0000-1000-8000-0026BB765291"

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
