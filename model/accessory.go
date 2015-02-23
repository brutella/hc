package model

// Base interface for all accessories
type Accessory interface {
	Compareable

	// GetId returns the accessory's id
	GetId() int64

	// GetServices returns the services which represent the accessory
	GetServices() []Service

	// Name returns accessory's name
	Name() string

	// SerialNumber returns the accessory's serial number
	SerialNumber() string

	// Manufacturer returns the accessory's manufacturer name
	Manufacturer() string

	// Model returns the accessory's model description
	Model() string

	// Firmware returns the accessory's firmware revision or empty string
	Firmware() string

	// Hardware returns the accessory's the hardware revision or empty string
	Hardware() string

	// Software returns the accessory's the sofware revision or empty string
	Software() string

	// OnIdentify calls the argument function to identify the accessory
	// Make the accessory identify itself (lights would blink)
	OnIdentify(func())
}
