package service

import (
	"github.com/brutella/hc/model"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	s := New()

	assert.Equal(t, s.GetID(), model.InvalidID)
	assert.Equal(t, len(s.GetCharacteristics()), 0)
}
