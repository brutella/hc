package model

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestBoolCharacteristic(t *testing.T) {
    b := NewBoolCharacteristic(true)
    assert.True(t, b.Bool())
    b.SetBool(false)
    assert.False(t, b.Bool())
}