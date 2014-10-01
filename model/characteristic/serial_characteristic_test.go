package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestSerialCharacteristic(t *testing.T) {
    str := NewSerialNumberCharacteristic("001002")
    assert.Equal(t, str.Type, CharTypeSerialNumber)
    assert.Equal(t, str.SerialNumber(), "001002")
    str.SetSerialNumber("001003")
    assert.Equal(t, str.SerialNumber(), "001003")
}