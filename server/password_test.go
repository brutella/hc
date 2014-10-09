package server

import (            
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
    pwd, err := NewPassword("00011222")
    assert.Nil(t, err)
    assert.Equal(t, pwd, "000-11-222")
}

func TestShortPassword(t *testing.T) {
    _, err := NewPassword("0001122")
    assert.NotNil(t, err)
}

func TestLongPassword(t *testing.T) {
    _, err := NewPassword("000112221")
    assert.NotNil(t, err)
}
