package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseId(t *testing.T) {
	aid, cid, err := ParseAccessoryAndCharacterId("3.9")
	assert.Nil(t, err)
	assert.Equal(t, aid, 3)
	assert.Equal(t, cid, 9)
}

func TestParseInvalidId(t *testing.T) {
	_, _, err := ParseAccessoryAndCharacterId("random")
	assert.NotNil(t, err)
}
