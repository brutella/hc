package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestInUse(t *testing.T) {
    use := NewInUse(true)
    assert.Equal(t, use.Permissions, PermsRead())
    assert.True(t, use.InUse())
    use.SetInUse(false)
    assert.False(t, use.InUse())
}