// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	SwingModeSwingDisabled int = 0
	SwingModeSwingEnabled  int = 1
)

const TypeSwingMode = "B6"

type SwingMode struct {
	*Int
}

func NewSwingMode() *SwingMode {
	char := NewInt(TypeSwingMode)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &SwingMode{char}
}
