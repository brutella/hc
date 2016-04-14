// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeHardwareRevision = "53"

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
