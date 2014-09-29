package hk

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestByteCharacteristic(t *testing.T) {
    b := NewByteCharacteristic(0xFA)
    assert.Equal(t, b.Byte(), byte(0xFA))
    b.SetByte(0xAF)
    assert.Equal(t, b.Byte(), byte(0xAF))
}