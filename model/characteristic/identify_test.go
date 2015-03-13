package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteOnlyIdentifyCharacteristic(t *testing.T) {
	i := NewIdentify()
	assert.Equal(t, i.Type, CharTypeIdentify)
	assert.Nil(t, i.GetValue())
	i.SetBool(true)
	assert.Nil(t, i.GetValue())
	i.SetValueFromRemote(true)
	assert.Nil(t, i.GetValue())
	i.SetValue(true)
	assert.Nil(t, i.GetValue())
}

func TestWriteOnlyCharacteristicRemoteDelegate(t *testing.T) {
	c := NewIdentify()

	var oldValue interface{}
	var newValue interface{}
	c.OnRemoteChange(func(c *Characteristic, new, old interface{}) {
		newValue = new
		oldValue = old
	})

	c.SetValueFromRemote(true)
	assert.Equal(t, oldValue, nil)
	assert.Equal(t, newValue, true)
	assert.Nil(t, c.GetValue())
	c.SetValueFromRemote(false)
	assert.Equal(t, oldValue, nil)
	assert.Equal(t, newValue, false)
	assert.Nil(t, c.GetValue())
}
