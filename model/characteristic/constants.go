package characteristic

import (
	"github.com/brutella/hc/model"
)

const (

	// Encoded in HomeKit Accessory Simulator.app as bits
	// 2:  0010 -> write
	// 15: 1111 -> read, write, events, bonjour
	// 13: 1101 -> read, events, bonjour

	PermRead   = "pr" // can be read
	PermWrite  = "pw" // can be written
	PermEvents = "ev" // sends events

	// Unused
	PermBonjour = "bonjour"
)

// PermsAll returns read, write and event permissions
func PermsAll() []string {
	return []string{PermRead, PermWrite, PermEvents}
}

// PermsRead returns read and event permissions
func PermsRead() []string {
	return []string{PermRead, PermEvents}
}

// PermsWrite returns write and event permissions
func PermsWrite() []string {
	return []string{PermWrite, PermEvents}
}

// PermsReadOnly returns read permission
func PermsReadOnly() []string {
	return []string{PermRead}
}

// PermsWriteOnly returns write permission
func PermsWriteOnly() []string {
	return []string{PermWrite}
}

const (
	TempUnitCelsiusByte    = 0x00
	TempUnitFahrenheitByte = 0x01

	UnitPercentage = "percentage"
	// UnitPercentByte = 0x01 // TODO not sure
	UnitArcDegrees = "arcdegrees"
	// UnitArcDegreesByte = 0x02 // TODO not sure
)

// ByteFromTempUnit returns the byte representing the argument TempUnit.
func ByteFromTempUnit(unit model.TempUnit) byte {
	switch unit {
	case model.TempUnitCelsius:
		return TempUnitCelsiusByte
	case model.TempUnitFahrenheit:
		return TempUnitFahrenheitByte
	}

	return 0x00
}

// TempUnitFromByte returns the TempUnit representing the byte.
func TempUnitFromByte(b byte) model.TempUnit {
	switch b {
	case TempUnitCelsiusByte:
		return model.TempUnitCelsius
	}

	return "Unknown"
}

// HomeKit defined charactersitic value types
const (
	FormatString = "string" // maxLen appears
	FormatBool   = "bool"   // on|off
	FormatInt    = "int"    // minValue, maxValue and minStep appear
	FormatFloat  = "float"  // minValue, maxValue, minStep and precision appear
	FormatByte   = "uint8"
	FormatTLV8   = "tlv8"
)

type CharType string

// HomeKit defined charactersitic types
const (
	CharTypeUnknown CharType = "0" // not specified in HAP

	// name service
	CharTypeName CharType = "23" // 13

	// info service
	CharTypeModel        CharType = "21" // 1
	CharTypeManufacturer CharType = "20" // 1
	CharTypeSerialNumber CharType = "30" // 1
	CharTypeIdentify     CharType = "14" // 2

	CharTypeLogs                     CharType = "1F" // 15
	CharTypeVersion                  CharType = "37" // 2
	CharTypeAdministratorOnlyAccesss CharType = "1"  // 15
	CharTypeFirmwareRevision         CharType = "52" // 13?
	CharTypeHardwareRevision         CharType = "53" // 13?
	CharTypeSoftwareRevision         CharType = "54" // 13?

	// Light bulb service
	CharTypeBrightness CharType = "8"  // 15
	CharTypeHue        CharType = "13" // 15
	CharTypeSaturation CharType = "2F" // 15

	// switch/outlet service
	CharTypePowerState CharType = "25" // 15

	CharTypeInUse         CharType = "26" // 13
	CharTypeAudioFeedback CharType = "5"  // 15

	// garage door opener
	CharTypeObstructionDetected                  CharType = "24" // 13
	CharTypeDoorStateTarget                      CharType = "32" // 15
	CharTypeDoorStateCurrent                     CharType = "E"  // 13
	CharTypeLockMechanismTargetState             CharType = "1E" // 15
	CharTypeLockMechanismCurrentState            CharType = "1D" // 13
	CharTypeLockMechanismLastKnownAction         CharType = "1C" // 13
	CharTypeLockMechanismAdditionalAuthorization CharType = "1B"
	CharTypeLockManagementControlPoint           CharType = "19" // 2
	CharTypeLockManagementAutoSecureTimeout      CharType = "1A"

	CharTypeRotationDirection CharType = "28"
	CharTypeRotationSpeed     CharType = "29"

	CharTypeTemperatureUnits   CharType = "36" // 15
	CharTypeTemperatureTarget  CharType = "35" // 15
	CharTypeTemperatureCurrent CharType = "11" // 13

	CharTypeRelativeHumidityTarget  CharType = "34" // 15
	CharTypeRelativeHumidityCurrent CharType = "10" // 13

	CharTypeHeatingThreshold          CharType = "12" // 15
	CharTypeCoolingThreshold          CharType = "D"  // 15
	CharTypeHeatingCoolingModeTarget  CharType = "33" // 15
	CharTypeHeatingCoolingModeCurrent CharType = "F"  // 13

	CharTypeMotionDetected CharType = "22" // 13
)
