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

type CharacteristicType string

// HomeKit defined charactersitic types
const (
	TypeUnknown CharacteristicType = "0" // not specified in HAP

	// name service
	TypeName CharacteristicType = "23" // 13

	// info service
	TypeModel        CharacteristicType = "21" // 1
	TypeManufacturer CharacteristicType = "20" // 1
	TypeSerialNumber CharacteristicType = "30" // 1
	TypeIdentify     CharacteristicType = "14" // 2

	TypeLogs                     CharacteristicType = "1F" // 15
	TypeVersion                  CharacteristicType = "37" // 2
	TypeAdministratorOnlyAccesss CharacteristicType = "1"  // 15
	TypeFirmwareRevision         CharacteristicType = "52" // 13?
	TypeHardwareRevision         CharacteristicType = "53" // 13?
	TypeSoftwareRevision         CharacteristicType = "54" // 13?

	// Light bulb service
	TypeBrightness CharacteristicType = "8"  // 15
	TypeHue        CharacteristicType = "13" // 15
	TypeSaturation CharacteristicType = "2F" // 15

	// switch/outlet service
	TypePowerState CharacteristicType = "25" // 15

	TypeInUse         CharacteristicType = "26" // 13
	TypeAudioFeedback CharacteristicType = "5"  // 15

	// garage door opener
	TypeObstructionDetected                  CharacteristicType = "24" // 13
	TypeDoorStateTarget                      CharacteristicType = "32" // 15
	TypeDoorStateCurrent                     CharacteristicType = "E"  // 13
	TypeLockMechanismTargetState             CharacteristicType = "1E" // 15
	TypeLockMechanismCurrentState            CharacteristicType = "1D" // 13
	TypeLockMechanismLastKnownAction         CharacteristicType = "1C" // 13
	TypeLockMechanismAdditionalAuthorization CharacteristicType = "1B"
	TypeLockManagementControlPoint           CharacteristicType = "19" // 2
	TypeLockManagementAutoSecureTimeout      CharacteristicType = "1A"

	TypeRotationDirection CharacteristicType = "28"
	TypeRotationSpeed     CharacteristicType = "29"

	TypeTemperatureUnits   CharacteristicType = "36" // 15
	TypeTemperatureTarget  CharacteristicType = "35" // 15
	TypeTemperatureCurrent CharacteristicType = "11" // 13

	TypeRelativeHumidityTarget  CharacteristicType = "34" // 15
	TypeRelativeHumidityCurrent CharacteristicType = "10" // 13

	TypeHeatingThreshold          CharacteristicType = "12" // 15
	TypeCoolingThreshold          CharacteristicType = "D"  // 15
	TypeHeatingCoolingModeTarget  CharacteristicType = "33" // 15
	TypeHeatingCoolingModeCurrent CharacteristicType = "F"  // 13

	TypeMotionDetected CharacteristicType = "22" // 13
)
