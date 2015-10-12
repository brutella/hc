package event

// DevicePaired is emitted when transport paired with a device (e.g. iOS client successfully paired with the accessory)
type DevicePaired struct{}

// DeviceUnpaired is emitted when pairing with a device is removed (e.g. iOS client removed the accessory from HomeKit)
type DeviceUnpaired struct{}
