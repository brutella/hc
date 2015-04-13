package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialForName(t *testing.T) {
	storage, err := NewTempFileStorage()
	assert.Nil(t, err)
	name := "My Accessory"
	serial := GetSerialNumberForAccessoryName(name, storage)
	serial2 := GetSerialNumberForAccessoryName(name, storage)
	assert.Equal(t, serial, serial2)
}
