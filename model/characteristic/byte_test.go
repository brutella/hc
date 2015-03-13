package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteCharacteristic(t *testing.T) {
	b := NewByteCharacteristic(0xFA, PermsAll())
	assert.Equal(t, b.Byte(), byte(0xFA))
	b.SetByte(0xAF)
	assert.Equal(t, b.Byte(), byte(0xAF))
}
