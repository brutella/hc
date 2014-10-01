package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
    b := NewBool(true)
    assert.True(t, b.BoolValue())
    b.SetBool(false)
    assert.False(t, b.BoolValue())
}