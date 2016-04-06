// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeHardwareRevision = "00000053-0000-1000-8000-0026BB765291"

type HardwareRevision struct {
	*String
}

func NewHardwareRevision() *HardwareRevision {
	char := NewString(TypeHardwareRevision)
	char.Format = FormatString
	char.Perms = []string{PermRead}

	char.SetValue("")

	return &HardwareRevision{char}
}
