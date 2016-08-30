// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ChargingStateNotCharging   int = 0
	ChargingStateCharging      int = 1
	ChargingStateNotChargeable int = 2
)

const TypeChargingState = "8F"

type ChargingState struct {
	*Int
}

func NewChargingState() *ChargingState {
	char := NewInt(TypeChargingState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &ChargingState{char}
}
