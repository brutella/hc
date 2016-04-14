// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeFirmwareRevision = "52"

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
