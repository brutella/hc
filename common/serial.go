package common

// SerialFilenameForName returns the serial file name for a name
func serialFilenameForName(name string) string {
	return name + ".serial"
}

// GetSerialNumberForAccessoryName returns the serial for a specific name stored in storage.
// When no serial number is stored for this name yet, a new one is created
// using RandomHexString()
func GetSerialNumberForAccessoryName(name string, storage Storage) string {
	serial_file := serialFilenameForName(name)
	serial_bytes, _ := storage.Get(serial_file)
	serial := string(serial_bytes)
	if len(serial) == 0 {
		serial = RandomHexString()
		storage.Set(serial_file, []byte(serial))
	}

	return serial
}
