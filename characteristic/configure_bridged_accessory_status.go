// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeConfigureBridgedAccessoryStatus = "9D"

type ConfigureBridgedAccessoryStatus struct {
	*Bytes
}

func NewConfigureBridgedAccessoryStatus() *ConfigureBridgedAccessoryStatus {
	char := NewBytes(TypeConfigureBridgedAccessoryStatus)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue([]byte{})

	return &ConfigureBridgedAccessoryStatus{char}
}
