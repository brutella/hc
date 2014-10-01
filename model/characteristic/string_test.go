package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
    str := NewString("A String")
    assert.Equal(t, str.StringValue(), "A String")
    str.SetString("My String")
    assert.Equal(t, str.StringValue(), "My String")
}