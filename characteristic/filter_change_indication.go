// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	FilterChangeIndicationFilterOK     int = 0
	FilterChangeIndicationChangeFilter int = 1
)

const TypeFilterChangeIndication = "AC"

type FilterChangeIndication struct {
	*Int
}

func NewFilterChangeIndication() *FilterChangeIndication {
	char := NewInt(TypeFilterChangeIndication)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &FilterChangeIndication{char}
}
