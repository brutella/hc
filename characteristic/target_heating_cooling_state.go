// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetHeatingCoolingStateOff  int = 0
	TargetHeatingCoolingStateHeat int = 1
	TargetHeatingCoolingStateCool int = 2
	TargetHeatingCoolingStateAuto int = 3
)

const TypeTargetHeatingCoolingState = "33"

type TargetHeatingCoolingState struct {
	*Int
}

func NewTargetHeatingCoolingState() *TargetHeatingCoolingState {
	char := NewInt(TypeTargetHeatingCoolingState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &TargetHeatingCoolingState{char}
}
