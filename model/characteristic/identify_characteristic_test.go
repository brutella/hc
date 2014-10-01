package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestIdentifyCharacteristic(t *testing.T) {
    i := NewIdentifyCharacteristic(true)
    assert.Equal(t, i.Type, CharTypeIdentify)
    assert.True(t, i.Identify())
    i.SetIdentify(false)
    assert.False(t, i.Identify())
}