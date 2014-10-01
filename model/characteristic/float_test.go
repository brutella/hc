package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
    float := NewFloat(20.2)
    assert.Equal(t, float.FloatValue(), 20.2)
    float.SetFloat(10.1)
    assert.Equal(t, float.FloatValue(), 10.1)
}