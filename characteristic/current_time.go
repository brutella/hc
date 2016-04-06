// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentTime = "0000009B-0000-1000-8000-0026BB765291"

type CurrentTime struct {
	*String
}

func NewCurrentTime() *CurrentTime {
	char := NewString(TypeCurrentTime)
	char.Format = FormatString
	char.Perms = []string{PermRead, PermWrite}

	char.SetValue("")

	return &CurrentTime{char}
}
