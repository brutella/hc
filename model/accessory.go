package model

// Base interface for all accessories
type Accessory interface {
	Compareable

	// Returns the accessory id
	GetId() int64

	// Returns the services which represent the accessory
	GetServices() []Service

	// Returns the name
	Name() string

	// Returns the serial number
	SerialNumber() string

	// Returns the manufacturer name
	Manufacturer() string

	// Returns the model description
	Model() string

	// Returns the firmware revision or empty string
	Firmware() string

	// Returns the hardware revision or empty string
	Hardware() string

	// Returns the sofware revision or empty string
	Software() string

	// Callback to identify accessory
	// Make the accessory identify itself (lights would blink)
	OnIdentify(func())
}
