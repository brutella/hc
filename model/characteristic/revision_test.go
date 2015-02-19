package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardwareRevisionCharacteristic(t *testing.T) {
	hw := NewHardwareRevision("1.0")
	assert.Equal(t, hw.Type, CharTypeHardwareRevision)
	assert.Equal(t, hw.Revision(), "1.0")
	hw.SetRevision("1.0.1")
	assert.Equal(t, hw.Revision(), "1.0.1")
}

func TestFirmwareRevisionCharacteristic(t *testing.T) {
	hw := NewFirmwareRevision("1.0")
	assert.Equal(t, hw.Type, CharTypeFirmwareRevision)
}

func TestSoftwareRevisionCharacteristic(t *testing.T) {
	hw := NewSoftwareRevision("1.0")
	assert.Equal(t, hw.Type, CharTypeSoftwareRevision)
}
