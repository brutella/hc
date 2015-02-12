package accessory

import (
	"github.com/brutella/hap/model"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessoryIdentifyChanged(t *testing.T) {
	a := New(info)

	var identifyCalled = 0
	a.OnIdentify(func() {
		identifyCalled += 1
	})

	a.Info.Identify.SetValueFromRemote(true)
	// Identify is set to false immediately
	assert.False(t, a.Info.Identify.Identify())
	assert.Equal(t, identifyCalled, 1)
}
