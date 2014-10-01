package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestManufacturer(t *testing.T) {
    m := NewManufacturer("Apple")
    assert.Equal(t, m.Type, CharTypeManufacturer)
    assert.Equal(t, m.Manufacturer(), "Apple")
    m.SetManufacturer("Google")
    assert.Equal(t, m.Manufacturer(), "Google")
}