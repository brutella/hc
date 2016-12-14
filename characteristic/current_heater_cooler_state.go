// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentHeaterCoolerStateInactive int = 0
	CurrentHeaterCoolerStateIdle     int = 1
	CurrentHeaterCoolerStateHeating  int = 2
	CurrentHeaterCoolerStateCooling  int = 3
)

const TypeCurrentHeaterCoolerState = "B1"

type CurrentHeaterCoolerState struct {
	*Int
}

func NewCurrentHeaterCoolerState() *CurrentHeaterCoolerState {
	char := NewInt(TypeCurrentHeaterCoolerState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CurrentHeaterCoolerState{char}
}
