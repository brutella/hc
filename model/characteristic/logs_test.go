package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestLogs(t *testing.T) {
    n := NewLogs("Test")
    assert.Equal(t, n.Type, CharTypeLogs)
}