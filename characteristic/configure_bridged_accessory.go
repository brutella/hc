// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeConfigureBridgedAccessory = "A0"

type ConfigureBridgedAccessory struct {
	*Bytes
}

func NewConfigureBridgedAccessory() *ConfigureBridgedAccessory {
	char := NewBytes(TypeConfigureBridgedAccessory)
	char.Format = FormatTLV8
	char.Perms = []string{PermWrite}

	return &ConfigureBridgedAccessory{char}
}
