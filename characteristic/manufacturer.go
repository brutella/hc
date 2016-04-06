// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeManufacturer = "00000020-0000-1000-8000-0026BB765291"

type Manufacturer struct {
	*String
}

func NewManufacturer() *Manufacturer {
	char := NewString(TypeManufacturer)
	char.Format = FormatString
	char.Perms = []string{PermRead}

	char.SetValue("")

	return &Manufacturer{char}
}
