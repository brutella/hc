package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogs(t *testing.T) {
	n := NewLogs("Test")
	assert.Equal(t, n.Type, CharTypeLogs)
}
