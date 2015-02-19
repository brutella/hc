package service

type ServiceType string

const (
	TypeAccessoryInfo    = ServiceType("3E")
	TypeGarageDoorOpener = ServiceType("41")
	TypeLightBulb        = ServiceType("43")
	TypeLockManagement   = ServiceType("44")
	TypeLockMechanism    = ServiceType("45")
	TypeOutlet           = ServiceType("47")
	TypeSwitch           = ServiceType("49")
	TypeThermostat       = ServiceType("4A")
	TypeFan              = ServiceType("40")
)
