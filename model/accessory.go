package model

// The Accessory interface forms the basis for all HomeKit accessories.
type Accessory interface {
	Compareable

	// GetID returns the accessory id
	GetID() int64

	// GetServices returns the services which represent the accessory
	GetServices() []Service

	// Name returns accessory name
	Name() string

	// SerialNumber returns the accessory serial number
	SerialNumber() string

	// Manufacturer returns the accessory manufacturer name
	Manufacturer() string

	// Model returns the accessory model description
	Model() string

	// Firmware returns the accessory firmware revision or empty string
	Firmware() string

	// Hardware returns the accessory the hardware revision or empty string
	Hardware() string

	// Software returns the accessory the sofware revision or empty string
	Software() string

	// OnIdentify calls the argument function to identify the accessory
	// Make the accessory identify itself (lights would blink)
	OnIdentify(func())
}
