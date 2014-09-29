package hk

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestManufacturerCharacteristic(t *testing.T) {
    m := NewManufacturerCharacteristic("Apple")
    assert.Equal(t, m.Type, CharTypeManufacturer)
    assert.Equal(t, m.Manufacturer(), "Apple")
    m.SetManufacturer("Google")
    assert.Equal(t, m.Manufacturer(), "Google")
}