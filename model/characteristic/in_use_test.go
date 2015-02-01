package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInUse(t *testing.T) {
	use := NewInUse(true)
	assert.Equal(t, use.Permissions, PermsRead())
	assert.True(t, use.InUse())
	use.SetInUse(false)
	assert.False(t, use.InUse())
}
