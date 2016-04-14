// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeManufacturer = "20"

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
