// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	InUseNotInUse int = 0
	InUseInUse    int = 1
)

const TypeInUse = "D2"

type InUse struct {
	*Int
}

func NewInUse() *InUse {
	char := NewInt(TypeInUse)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &InUse{char}
}
