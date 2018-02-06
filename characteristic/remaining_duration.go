// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeRemainingDuration = "D4"

type RemainingDuration struct {
	*Int
}

func NewRemainingDuration() *RemainingDuration {
	char := NewInt(TypeRemainingDuration)
	char.Format = FormatUInt32
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(3600)
	char.SetStepValue(1)
	char.SetValue(0)

	return &RemainingDuration{char}
}
