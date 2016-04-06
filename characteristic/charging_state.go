// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ChargingStateNotCharging int = 0
	ChargingStateCharging    int = 1
)

const TypeChargingState = "0000008F-0000-1000-8000-0026BB765291"

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
