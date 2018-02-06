// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	IsConfiguredNotConfigured int = 0
	IsConfiguredConfigured    int = 1
)

const TypeIsConfigured = "D6"

type IsConfigured struct {
	*Int
}

func NewIsConfigured() *IsConfigured {
	char := NewInt(TypeIsConfigured)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &IsConfigured{char}
}
