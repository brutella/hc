package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestStringCharacteristic(t *testing.T) {
    str := NewStringCharacteristic("A String")
    assert.Equal(t, str.String(), "A String")
    str.SetString("My String")
    assert.Equal(t, str.String(), "My String")
}