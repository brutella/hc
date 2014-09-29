package model

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestFloatCharacteristic(t *testing.T) {
    float := NewFloatCharacteristic(20.2)
    assert.Equal(t, float.Float(), 20.2)
    float.SetFloat(10.1)
    assert.Equal(t, float.Float(), 10.1)
}