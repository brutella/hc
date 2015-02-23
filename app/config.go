package app

// Config contains essential information to setup and publish a HomeKit bridge
type Config struct {
	// Name of the bridge which appears in the HomeKit accessory browser on iOS
	BridgeName string
	// Password the user has to enter when adding the accessory to HomeKit
	BridgePassword string
	// Manufacturer name which appears in the bridge's accessory info service
	BridgeManufacturer string
	// Path to database folder
	DatabaseDir string
}

func NewConfig() Config {
	return Config{
		BridgeName:         "GoBridge",
		BridgePassword:     "001-02-003",
		BridgeManufacturer: "brutella",
	}
}
