package characteristic

const (
	PermRead   = "pr" // can be read
	PermWrite  = "pw" // can be written
	PermEvents = "ev" // sends events
)

// PermsAll returns read, write and event permissions
func PermsAll() []string {
	return []string{PermRead, PermWrite, PermEvents}
}

// PermsRead returns read and event permissions
func PermsRead() []string {
	return []string{PermRead, PermEvents}
}

// PermsReadOnly returns read permission
func PermsReadOnly() []string {
	return []string{PermRead}
}

// PermsWriteOnly returns write permission
func PermsWriteOnly() []string {
	return []string{PermWrite}
}

// HAP characteristic units
const (
	UnitPercentage = "percentage"
	UnitArcDegrees = "arcdegrees"
	UnitCelsius    = "celsius"
)

// HAP characterisitic formats
const (
	FormatString = "string"
	FormatBool   = "bool"
	FormatFloat  = "float"
	FormatUInt8  = "uint8"
	FormatUInt16 = "uint16"
	FormatUInt32 = "uint32"
	FormatInt32  = "int32"
	FormatUInt64 = "uint64"
	FormatData   = "data"
	FormatTLV8   = "tlv8"
)
