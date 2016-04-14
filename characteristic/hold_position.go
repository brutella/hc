// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeHoldPosition = "6F"

type HoldPosition struct {
	*Bool
}

func NewHoldPosition() *HoldPosition {
	char := NewBool(TypeHoldPosition)
	char.Format = FormatBool
	char.Perms = []string{PermWrite}

	return &HoldPosition{char}
}
