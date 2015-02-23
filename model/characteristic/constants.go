package characteristic

import (
	"github.com/brutella/hap/model"
)

const (

	// Encoded in HomeKit Accessory Simulator.app as bits
	// 2:  0010 -> write
	// 15: 1111 -> read, write, events, bonjour
	// 13: 1101 -> read, events, bonjour

	PermRead  = "pr" // can be read
	PermWrite = "pw" // can be written

	// Unused
	PermEvents  = "ev"
	PermBonjour = "bonjour" // TODO never used by accessory simulator
)

func PermsAll() []string {
	return []string{PermRead, PermWrite, PermEvents}
}

func PermsRead() []string {
	return []string{PermRead, PermEvents}
}

func PermsWrite() []string {
	return []string{PermWrite, PermEvents}
}

func PermsReadOnly() []string {
	return []string{PermRead}
}

func PermsWriteOnly() []string {
	return []string{PermWrite}
}

const (
	TempUnitCelsiusByte = 0x00

	UnitPercentage = "percentage"
	// UnitPercentByte = 0x01 // TODO not sure
	UnitArcDegrees = "arcdegrees"
	// UnitArcDegreesByte = 0x02 // TODO not sure
)

func ByteFromUnit(unit model.TempUnit) byte {
	switch unit {
	case model.TempUnitCelsius:
		return TempUnitCelsiusByte
	}

	return 0x00
}

func TempUnitFromByte(b byte) model.TempUnit {
	switch b {
	case TempUnitCelsiusByte:
		return model.TempUnitCelsius
	}

	return "Unknown"
}

const (
	FormatString = "string" // maxLen appears
	FormatBool   = "bool"   // on|off
	FormatInt    = "int"    // minValue, maxValue and minStep appear
	FormatFloat  = "float"  // minValue, maxValue, minStep and precision appear
	FormatByte   = "uint8"
)

type CharType string

const (
	CharTypeUnknown = "0" // not specified in HAP

	// name service
	CharTypeName = "23" // 13

	// info service
	CharTypeModel        = "21" // 1
	CharTypeManufacturer = "20" // 1
	CharTypeSerialNumber = "30" // 1
	CharTypeIdentify     = "14" // 2

	CharTypeLogs                     = "1F" // 15
	CharTypeVersion                  = "37" // 2
	CharTypeAdministratorOnlyAccesss = "1"  // 15
	CharTypeFirmwareRevision         = "52" // 13?
	CharTypeHardwareRevision         = "53" // 13?
	CharTypeSoftwareRevision         = "54" // 13?

	// Light bulb service
	CharTypeBrightness = "8"  // 15
	CharTypeHue        = "13" // 15
	CharTypeSaturation = "2F" // 15

	// switch/outlet service
	CharTypePowerState = "25" // 15

	CharTypeInUse         = "26" // 13
	CharTypeAudioFeedback = "5"  // 15

	// garage door opener
	CharTypeObstructionDetected                  = "24" // 13
	CharTypeDoorStateTarget                      = "32" // 15
	CharTypeDoorStateCurrent                     = "E"  // 13
	CharTypeLockMechanismTargetState             = "1E" // 15
	CharTypeLockMechanismCurrentState            = "1D" // 13
	CharTypeLockMechanismLastKnownAction         = "1C" // 13
	CharTypeLockMechanismAdditionalAuthorization = "1B"
	CharTypeLockManagementControlPoint           = "19" // 2
	CharTypeLockManagementAutoSecureTimeout      = "1A"

	CharTypeRotationDirection = "28"
	CharTypeRotationSpeed     = "29"

	CharTypeTemperatureUnits   = "36" // 15
	CharTypeTemperatureTarget  = "35" // 15
	CharTypeTemperatureCurrent = "11" // 13

	CharTypeRelativeHumidityTarget  = "34" // 15
	CharTypeRelativeHumidityCurrent = "10" // 13

	CharTypeHeatingThreshold          = "12" // 15
	CharTypeCoolingThreshold          = "D"  // 15
	CharTypeHeatingCoolingModeTarget  = "33" // 15
	CharTypeHeatingCoolingModeCurrent = "F"  // 13

	CharTypeMotionDetected = "22" // 13
)
