package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestOn(t *testing.T) {
    b := NewOn(true)
    assert.True(t, b.On())
    b.SetOn(false)
    assert.False(t, b.On())
}