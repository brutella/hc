package service

import (
    "github.com/brutella/hap/model"
    
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {    
     s := New()
    
    assert.Equal(t, s.GetId(), model.InvalidId)
    assert.Equal(t, len(s.GetCharacteristics()), 0)
}