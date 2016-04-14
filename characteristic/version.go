// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeVersion = "37"

type Version struct {
	*String
}

func NewVersion() *Version {
	char := NewString(TypeVersion)
	char.Format = FormatString
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue("")

	return &Version{char}
}
