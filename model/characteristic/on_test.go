package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOn(t *testing.T) {
	b := NewOn(true)
	assert.True(t, b.On())
	b.SetOn(false)
	assert.False(t, b.On())
}
