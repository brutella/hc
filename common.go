package hap

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

const (
    CharTypeName                            = 0x00000023
    CharTypeModel                           = 0x00000021
    CharTypeManufacturer                    = 0x00000020
    CharTypeSerialNumber                    = 0x00000030
    CharTypeLogs                            = 0x00000025
    CharTypeVersion                         = 0x00000037
    CharTypeIdentify                        = 0x00000014
    CharTypeAdministratorOnlyAccesss        = 0x00000001
    
    CharTypeBrightness                      = 0x00000008
    CharTypeHue                             = 0x00000013
    CharTypeSaturation                      = 0x0000002F
    
    CharTypeOn                              = 0x00000025 // to identify
    CharTypeOutletInUse                     = 0x00000026
    CharTypeAudioFeedback                   = 0x00000005
    
    CharTypeDoorStateTarget                 = 0x00000032
    CharTypeDoorStateCurrent                = 0x0000000E
    
    CharTypeTemperatureUnits                = 0x00000036
    CharTypeTemperatureTarget               = 0x00000035
    CharTypeTemperatureCurrent              = 0x00000011
    
    CharTypeRelativeHumidityTarget          = 0x00000034
    CharTypeRelativeHumidityCurrent         = 0x00000010
    
    CharTypeHeatingThreshold                = 0x00000012
    CharTypeHeatingCoolingModeTarget        = 0x00000033
    CharTypeHeatingCoolingModeCurrent       = 0x0000000F

    CharTypeLockMechanismTargetState        = 0x0000001E
    CharTypeLockMechanismLastKnownAction    = 0x0000001C
    CharTypeLockMechanismCurrentState       = 0x0000001D
    CharTypeLockManagementControlPoint      = 0x00000019
    
    CharTypeObstructionDetected             = 0x00000024
    CharTypeMotionDetected                  = 0x00000022
)

const (
    ServiceTypeAccessoryInfo                = 0x0000003E
    ServiceTypeGarageDoorOpener             = 0x00000041
    ServiceTypeLightBulb                    = 0x00000043
    ServiceTypeLockManagement               = 0x00000044
    ServiceTypeLockMechanism                = 0x00000045
    ServiceTypeOutlet                       = 0x00000047
    ServiceTypeSwitch                       = 0x00000049
    SerivceTypeThermostat                   = 0x0000004A
)