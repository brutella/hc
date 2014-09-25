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