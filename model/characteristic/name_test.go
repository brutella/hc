package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
    n := NewName("Test")
    assert.Equal(t, n.Type, CharTypeName)
    assert.Equal(t, n.Name(), "Test")
    n.SetName("My Name")
    assert.Equal(t, n.Name(), "My Name")
}