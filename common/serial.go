package common

// SerialFilenameForName returns the serial file name for a name
func serialFilenameForName(name string) string {
	return name + ".serial"
}

// GetSerialNumberForAccessoryName returns the serial for a specific name stored in storage.
// When no serial number is stored for this name yet, a new one is created
// using RandomHexString()
func GetSerialNumberForAccessoryName(name string, storage Storage) string {
	filename := serialFilenameForName(name)
	b, _ := storage.Get(filename)
	str := string(b)
	if len(str) == 0 {
		str = RandomHexString()
		storage.Set(filename, []byte(str))
	}

	return str
}
