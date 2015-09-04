package service

// serviceType is the type for all HomeKit service types.
type serviceType string

// HomeKit defined service types.
const (
	typeAccessoryInfo               serviceType = "3E"
	typeAirQualitySensor            serviceType = "8D"
	typeBatteryService              serviceType = "96"
	typeBridgingState               serviceType = "62"
	typeCarbonDioxideSensor         serviceType = "97"
	typeCarbonMonoxideSensor        serviceType = "7F"
	typeContactSensor               serviceType = "80"
	typeDoor                        serviceType = "81"
	typeFan                         serviceType = "40"
	typeGarageDoorOpener            serviceType = "41"
	typeHumiditySensor              serviceType = "82"
	typeLeakSensor                  serviceType = "83"
	typeLightSensor                 serviceType = "84"
	typeLightbulb                   serviceType = "43"
	typeLockManagement              serviceType = "44"
	typeLockMechanism               serviceType = "45"
	typeMotionSensor                serviceType = "85"
	typeOccupancySensor             serviceType = "86"
	typeOutlet                      serviceType = "47"
	typeSecuritySystem              serviceType = "7E"
	typeSmokeSensor                 serviceType = "87"
	typeStatefulProgrammableSwitch  serviceType = "88"
	typeStatelessProgrammableSwitch serviceType = "89"
	typeSwitch                      serviceType = "49"
	typeTemperatureSensor           serviceType = "8A"
	typeThermostat                  serviceType = "4A"
	typeWindow                      serviceType = "8B"
	typeWindowCovering              serviceType = "8C"
)
