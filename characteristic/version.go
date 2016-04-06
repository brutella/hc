// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeVersion = "00000037-0000-1000-8000-0026BB765291"

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
