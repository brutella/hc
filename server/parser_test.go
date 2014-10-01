package server

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestParseID(t *testing.T) {
    aid, cid, err := ParseAccessoryAndCharacterId("3.9")
    assert.Nil(t, err)
    assert.Equal(t, aid, 3)
    assert.Equal(t, cid, 9)
}