// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	PowerModeSelectionShow int = 0
	PowerModeSelectionHide int = 1
)

const TypePowerModeSelection = "DF"

type PowerModeSelection struct {
	*Int
}

func NewPowerModeSelection() *PowerModeSelection {
	char := NewInt(TypePowerModeSelection)
	char.Format = FormatUInt8
	char.Perms = []string{PermWrite}
	char.SetMinValue(0)
	char.SetMaxValue(1)
	char.SetStepValue(1)

	return &PowerModeSelection{char}
}
