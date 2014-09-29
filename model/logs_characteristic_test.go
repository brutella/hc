package hk

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestLogsCharacteristic(t *testing.T) {
    n := NewLogsCharacteristic("Test")
    assert.Equal(t, n.Type, CharTypeLogs)
}