package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

type delegate struct {
    CharacteristicDelegate
    oldValue interface{}
    newValue interface{}
}

func (d *delegate) CharactericDidChangeValue(c *Characteristic, change CharacteristicChange) {
    d.oldValue = change.OldValue
    d.newValue = change.NewValue
}

func TestCharacteristicSetValuesOfWrongType(t *testing.T) {
    var value int = 5
    c := NewCharacteristic(value, FormatInt, CharTypeOn, nil)
    
    c.SetValue(float64(20.5))
    assert.Equal(t, c.Value, 20)
    
    c.SetValue("91")
    assert.Equal(t, c.Value, 91)
    
    c.SetValue(true)
    assert.Equal(t, c.Value, 1)
}

func TestCharacteristicLocalDelegate(t *testing.T) {
    c := NewCharacteristic(5, FormatInt, CharTypeOn, nil)
    
    d := &delegate{}
    c.AddLocalChangeDelegate(d)
    c.SetValue(10)
    assert.Equal(t, d.oldValue, 5)
    assert.Equal(t, d.newValue, 10)
    c.SetValueFromRemote(20)
    assert.Equal(t, d.oldValue, 5)
    assert.Equal(t, d.newValue, 10)
}

func TestCharacteristicRemoteDelegate(t *testing.T) {
    c := NewCharacteristic(5, FormatInt, CharTypeOn, nil)
    
    d := &delegate{}
    c.AddRemoteChangeDelegate(d)
    c.SetValueFromRemote(10)
    assert.Equal(t, d.oldValue, 5)
    assert.Equal(t, d.newValue, 10)
    c.SetValue(20)
    assert.Equal(t, d.oldValue, 5)
    assert.Equal(t, d.newValue, 10)
}

func TestRemoveDelegate(t *testing.T) {
    c := NewCharacteristic(5, FormatInt, CharTypeOn, nil)
    
    d := &delegate{}
    c.AddLocalChangeDelegate(d)
    c.RemoveDelegate(d)
    c.SetValueFromRemote(10)
    c.SetValue(20)
    assert.Nil(t, d.oldValue)
    assert.Nil(t, d.newValue)
}

func TestEqual(t *testing.T) {
   c1 := NewCharacteristic(5, FormatInt, CharTypeOn, nil)
   c2 := NewCharacteristic(5, FormatInt, CharTypeOn, nil) 
   assert.True(t, c1.Equal(c2))
}
