package service

import (
	"github.com/brutella/hc/model"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	s := New()

	assert.Equal(t, s.GetId(), model.InvalidId)
	assert.Equal(t, len(s.GetCharacteristics()), 0)
}
