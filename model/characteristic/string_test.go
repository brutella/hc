package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	str := NewString("A String")
	assert.Equal(t, str.StringValue(), "A String")
	str.SetString("My String")
	assert.Equal(t, str.StringValue(), "My String")
}
