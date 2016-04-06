// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeHoldPosition = "0000006F-0000-1000-8000-0026BB765291"

type HoldPosition struct {
	*Bool
}

func NewHoldPosition() *HoldPosition {
	char := NewBool(TypeHoldPosition)
	char.Format = FormatBool
	char.Perms = []string{PermWrite}

	return &HoldPosition{char}
}
