// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAccessoryFlags = "A6"

type AccessoryFlags struct {
	*Int
}

func NewAccessoryFlags() *AccessoryFlags {
	char := NewInt(TypeAccessoryFlags)
	char.Format = FormatUInt32
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &AccessoryFlags{char}
}
