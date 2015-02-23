package common

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// SerialFilenameForName returns the serial file name for a name
func SerialFilenameForName(name string) string {
	return name + ".serial"
}

// GetSerialNumberForAccessoryName returns the serial for a specific name.
// When no serial number is stored for this name yet, a new one is created
// using GenerateSerialNumber()
func GetSerialNumberForAccessoryName(name string, storage Storage) string {
	serial_file := SerialFilenameForName(name)
	serial_bytes, _ := storage.Get(serial_file)
	serial := string(serial_bytes)
	if len(serial) == 0 {
		serial = GenerateSerialNumber()
		storage.Set(serial_file, []byte(serial))
	}

	return serial
}

// GenerateSerialNumber generates a new serial number using the current time stamp
func GenerateSerialNumber() string {
	t := time.Now().Format(time.RFC3339Nano)

	h := md5.New()
	h.Write([]byte(t))
	result := h.Sum(nil)

	return hex.EncodeToString(result)
}
