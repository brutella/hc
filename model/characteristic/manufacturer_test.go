package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManufacturer(t *testing.T) {
	m := NewManufacturer("Apple")
	assert.Equal(t, m.Type, CharTypeManufacturer)
	assert.Equal(t, m.Manufacturer(), "Apple")
	m.SetManufacturer("Google")
	assert.Equal(t, m.Manufacturer(), "Google")
}
