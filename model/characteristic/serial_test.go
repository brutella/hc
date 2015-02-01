package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialCharacteristic(t *testing.T) {
	str := NewSerialNumber("001002")
	assert.Equal(t, str.Type, CharTypeSerialNumber)
	assert.Equal(t, str.SerialNumber(), "001002")
	str.SetSerialNumber("001003")
	assert.Equal(t, str.SerialNumber(), "001003")
}
