// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTargetPositionStatus = "7C"

type TargetPositionStatus struct {
	*Int
}

func NewTargetPositionStatus() *TargetPositionStatus {
	char := NewInt(TypeTargetPositionStatus)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(0)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &TargetPositionStatus{char}
}
