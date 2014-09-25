package model

type UUID string
const (
    PermRead = "pr" // can be read
    PermWrite = "pw" // can be written
    
    // Unused
    PermEvents = "events" // can be watched for notifications
    PermBonjour = "bonjour" // ?
)

const (
    FormatString    = "string"  // maxLen appears
    FormatBool      = "bool"    // on|off
    FormatInt       = "int"     // minValue, maxValue and minStep appear
    FormatFloat     = "float"   // minValue, maxValue, minStep and precision appear
)

const (
    UnitCelsius     = "celsius"
    UnitPercent     = "percent"
    UnitArcDegrees  = "arcdegrees"
)

type CharType string
const (
    CharTypeUnknown                         = "0" // not specified in HAP
    
    CharTypeName                            = "23"
    CharTypeModel                           = "21"
    CharTypeManufacturer                    = "20"
    CharTypeSerialNumber                    = "30"
    CharTypeLogs                            = "25"
    CharTypeVersion                         = "37"
    CharTypeIdentify                        = "14" // to identify
    CharTypeAdministratorOnlyAccesss        = "1"
    
    CharTypeBrightness                      = "8"
    CharTypeHue                             = "13"
    CharTypeSaturation                      = "2F"
    
    CharTypeOn                              = "25"
    CharTypeOutletInUse                     = "26"
    CharTypeAudioFeedback                   = "05"
    
    CharTypeDoorStateTarget                 = "32"
    CharTypeDoorStateCurrent                = "0E"
    
    CharTypeTemperatureUnits                = "36"
    CharTypeTemperatureTarget               = "35"
    CharTypeTemperatureCurrent              = "11"
    
    CharTypeRelativeHumidityTarget          = "34"
    CharTypeRelativeHumidityCurrent         = "10"
    
    CharTypeHeatingThreshold                = "12"
    CharTypeHeatingCoolingModeTarget        = "33"
    CharTypeHeatingCoolingModeCurrent       = "0F"

    CharTypeLockMechanismTargetState        = "1E"
    CharTypeLockMechanismLastKnownAction    = "1C"
    CharTypeLockMechanismCurrentState       = "1D"
    CharTypeLockManagementControlPoint      = "19"
    
    CharTypeObstructionDetected             = "24"
    CharTypeMotionDetected                  = "22"
)

type ServiceType string
const (
    ServiceTypeAccessoryInfo                = "3E"
    ServiceTypeGarageDoorOpener             = "41"
    ServiceTypeLightBulb                    = "43"
    ServiceTypeLockManagement               = "44"
    ServiceTypeLockMechanism                = "45"
    ServiceTypeOutlet                       = "47"
    ServiceTypeSwitch                       = "49"
    SerivceTypeThermostat                   = "4A"
)