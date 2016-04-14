// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAccessoryIdentifier = "57"

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
