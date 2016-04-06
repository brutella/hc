// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeConfigureBridgedAccessoryStatus = "0000009D-0000-1000-8000-0026BB765291"

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
