package server

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Tests the pairing key verification
func TestParseAccessoryAndCharacterId(t *testing.T) {
    aid, cid, err := ParseAccessoryAndCharacterId("10.1")
    assert.Nil(t, err)
    assert.Equal(t, aid, 10)
    assert.Equal(t, cid, 1)
}