// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeConfigureBridgedAccessory = "000000A0-0000-1000-8000-0026BB765291"

type ConfigureBridgedAccessory struct {
	*Bytes
}

func NewConfigureBridgedAccessory() *ConfigureBridgedAccessory {
	char := NewBytes(TypeConfigureBridgedAccessory)
	char.Format = FormatTLV8
	char.Perms = []string{PermWrite}

	return &ConfigureBridgedAccessory{char}
}
