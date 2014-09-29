package hk

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestNameCharacteristic(t *testing.T) {
    n := NewNameCharacteristic("Test")
    assert.Equal(t, n.Type, CharTypeName)
    assert.Equal(t, n.Name(), "Test")
    n.SetName("My Name")
    assert.Equal(t, n.Name(), "My Name")
}