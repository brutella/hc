// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeConfiguredName = "E3"

type ConfiguredName struct {
	*String
}

func NewConfiguredName() *ConfiguredName {
	char := NewString(TypeConfiguredName)
	char.Format = FormatString
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue("")

	return &ConfiguredName{char}
}
