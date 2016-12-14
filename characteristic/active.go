// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ActiveInactive int = 0
	ActiveActive   int = 1
)

const TypeActive = "B0"

type Active struct {
	*Int
}

func NewActive() *Active {
	char := NewInt(TypeActive)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &Active{char}
}
