package service

// serviceType is the type for all HomeKit service types.
type serviceType string

// HomeKit defined service types.
const (
	typeAccessoryInfo    = serviceType("3E")
	typeGarageDoorOpener = serviceType("41")
	typeLightBulb        = serviceType("43")
	typeLockManagement   = serviceType("44")
	typeLockMechanism    = serviceType("45")
	typeOutlet           = serviceType("47")
	typeSwitch           = serviceType("49")
	typeThermostat       = serviceType("4A")
	typeFan              = serviceType("40")
)
