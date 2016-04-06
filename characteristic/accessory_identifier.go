// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAccessoryIdentifier = "00000057-0000-1000-8000-0026BB765291"

type AccessoryIdentifier struct {
	*String
}

func NewAccessoryIdentifier() *AccessoryIdentifier {
	char := NewString(TypeAccessoryIdentifier)
	char.Format = FormatString
	char.Perms = []string{PermRead}

	char.SetValue("")

	return &AccessoryIdentifier{char}
}
