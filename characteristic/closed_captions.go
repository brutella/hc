// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ClosedCaptionsDisabled int = 0
	ClosedCaptionsEnabled  int = 1
)

const TypeClosedCaptions = "DD"

type ClosedCaptions struct {
	*Int
}

func NewClosedCaptions() *ClosedCaptions {
	char := NewInt(TypeClosedCaptions)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1)
	char.SetStepValue(1)
	char.SetValue(0)

	return &ClosedCaptions{char}
}
