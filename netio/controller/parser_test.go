package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseID(t *testing.T) {
	aid, cid, err := ParseAccessoryAndCharacterID("3.9")
	assert.Nil(t, err)
	assert.Equal(t, aid, 3)
	assert.Equal(t, cid, 9)
}

func TestParseInvalidID(t *testing.T) {
	_, _, err := ParseAccessoryAndCharacterID("random")
	assert.NotNil(t, err)
}
