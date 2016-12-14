// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	CurrentSlatStateFixed    int = 0
	CurrentSlatStateJammed   int = 1
	CurrentSlatStateSwinging int = 2
)

const TypeCurrentSlatState = "AA"

type CurrentSlatState struct {
	*Int
}

func NewCurrentSlatState() *CurrentSlatState {
	char := NewInt(TypeCurrentSlatState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &CurrentSlatState{char}
}
