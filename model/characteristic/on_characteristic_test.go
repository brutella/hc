package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestOnCharacteristic(t *testing.T) {
    b := NewOnCharacteristic(true)
    assert.True(t, b.On())
    b.SetOn(false)
    assert.False(t, b.On())
}