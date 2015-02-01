package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModel(t *testing.T) {
	m := NewModel("Late 2014")
	assert.Equal(t, m.Type, CharTypeModel)
	assert.Equal(t, m.Model(), "Late 2014")
	m.SetModel("Early 2015")
	assert.Equal(t, m.Model(), "Early 2015")
}
