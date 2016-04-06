// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeFirmwareRevision = "00000052-0000-1000-8000-0026BB765291"

type FirmwareRevision struct {
	*String
}

func NewFirmwareRevision() *FirmwareRevision {
	char := NewString(TypeFirmwareRevision)
	char.Format = FormatString
	char.Perms = []string{PermRead}

	char.SetValue("")

	return &FirmwareRevision{char}
}
