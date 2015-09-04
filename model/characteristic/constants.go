package characteristic

import (
	"github.com/brutella/hc/model"
)

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

	TypeAdministratorOnlyAccesss        CharacteristicType = "1"  // rwn
	TypeAirParticulateDensityType       CharacteristicType = "64" // rn
	TypeAirParticulateSize              CharacteristicType = "65" // rn
	TypeAirQuality                      CharacteristicType = "95" // rn
	TypeAudioFeedback                   CharacteristicType = "5"  // rwn
	TypeBatteryLevel                    CharacteristicType = "68" // rn
	TypeBrightness                      CharacteristicType = "8"  // rwn
	TypeCarbonDioxideDetected           CharacteristicType = "92" // rn
	TypeCarbonDioxideLevel              CharacteristicType = "93" // rn
	TypeCarbonDioxidePeakLevel          CharacteristicType = "94" // rn
	TypeCarbonMonoxideDetected          CharacteristicType = "69" // rn
	TypeCarbonMonoxideLevel             CharacteristicType = "90" // rn
	TypeChargingState                   CharacteristicType = "8F" // rn
	TypeContactSensorState              CharacteristicType = "6A" // rn
	TypeCoolingThreshold                CharacteristicType = "D"  // rwn
	TypeCurrentAmbientLightLevel        CharacteristicType = "6B" //rn
	TypeCurrentDoorState                CharacteristicType = "E"  // rn
	TypeCurrentHeatingCoolingMode       CharacteristicType = "F"  // rn
	TypeCurrentHorizontalTiltAngle      CharacteristicType = "6C" // rn
	TypeCurrentPosition                 CharacteristicType = "6D" // rn
	TypeCurrentRelativeHumidity         CharacteristicType = "10" // rn
	TypeCurrentTemperature              CharacteristicType = "11" // rn
	TypeCurrentVerticalTiltAngle        CharacteristicType = "6E" // rn
	TypeFirmwareRevision                CharacteristicType = "52" // rn
	TypeHardwareRevision                CharacteristicType = "53" // rn
	TypeHeatingThresholdTemperature     CharacteristicType = "12" // rwn
	TypeHoldPositiong                   CharacteristicType = "6F" // w
	TypeHue                             CharacteristicType = "13" // rwn
	TypeIdentify                        CharacteristicType = "14" // w
	TypeLeakDetected                    CharacteristicType = "70" // rn
	TypeLockControlPoint                CharacteristicType = "19" // w
	TypeLockMechanismCurrentState       CharacteristicType = "1D" // rn
	TypeLockMechanismLastKnownAction    CharacteristicType = "1C" // rn
	TypeLockManagementAutoSecureTimeout CharacteristicType = "1A" // rwn
	TypeLockMechanismTargetState        CharacteristicType = "1E" // rwn
	TypeLogs                            CharacteristicType = "1F" // rn
	TypeManufacturer                    CharacteristicType = "20" // r
	TypeModel                           CharacteristicType = "21" // r
	TypeMotionDetected                  CharacteristicType = "22" // rn
	TypeName                            CharacteristicType = "23" // rn
	TypeObstructionDetected             CharacteristicType = "24" // rn
	TypeOccupancyDetected               CharacteristicType = "71" // rn
	TypePowerState                      CharacteristicType = "25" // rwn
	TypeOutletInUse                     CharacteristicType = "26" // rn
	TypePositionState                   CharacteristicType = "72" // rn
	TypeProgrammableSwitchEvent         CharacteristicType = "73" // rn
	TypeProgrammableSwitchOutputState   CharacteristicType = "74" // rwn
	TypeReachable                       CharacteristicType = "63" // rn
	TypeRotationDirection               CharacteristicType = "28" // rwn
	TypeRotationSpeed                   CharacteristicType = "29" // rwn
	TypeSaturation                      CharacteristicType = "2F" // rwn
	TypeSecuritySystemAlarmType         CharacteristicType = "8E" // rn
	TypeSecuritySystemCurrentState      CharacteristicType = "66" // rn
	TypeSecuritySystemTargetState       CharacteristicType = "67" // rwn
	TypeSerialNumber                    CharacteristicType = "30" // r
	TypeSmokeDetected                   CharacteristicType = "76" // rn
	TypeSoftwareRevision                CharacteristicType = "54" // rn
	TypeStatusActive                    CharacteristicType = "75" // rn
	TypeStatusFault                     CharacteristicType = "77" // rn
	TypeStatusJammed                    CharacteristicType = "78" // rn
	TypeStatusLowBattery                CharacteristicType = "79" // rn
	TypeStatusTampered                  CharacteristicType = "7A" // rn
	TypeTargetDoorState                 CharacteristicType = "32" // rwn
	TypeTargetHeatingCoolingMode        CharacteristicType = "33" // rwn
	TypeTargetHorizontalTiltAngle       CharacteristicType = "7B" // rwn
	TypeTargetPosition                  CharacteristicType = "7C" // rwn
	TypeTargetRelativeHumidity          CharacteristicType = "34" // rwn
	TypeTargetTemperature               CharacteristicType = "35" // rwn
	TypeTargetVerticalTiltAngle         CharacteristicType = "7D" // rwn
	TypeTemperatureDisplayUnits         CharacteristicType = "36" // rwn
	TypeVersion                         CharacteristicType = "37" // rwn
)
