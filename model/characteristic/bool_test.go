package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	b := NewBool(true)
	assert.True(t, b.BoolValue())
	b.SetBool(false)
	assert.False(t, b.BoolValue())
}
