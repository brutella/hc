package container

import (
	"github.com/brutella/hap/model"
	"github.com/brutella/hap/model/accessory"

	"github.com/stretchr/testify/assert"
	"testing"
)

var info = model.Info{
	Name:         "Accessory1",
	SerialNumber: "001",
	Manufacturer: "Google",
	Model:        "Accessory",
}

func TestModel(t *testing.T) {
	acc1 := accessory.New(info)
	assert.Equal(t, acc1.GetId(), model.InvalidId)

	info.Name = "Accessory2"
	acc2 := accessory.New(info)
	assert.Equal(t, acc2.GetId(), model.InvalidId)

	c := NewContainer()
	c.AddAccessory(acc1)
	c.AddAccessory(acc2)

	assert.Equal(t, len(c.Accessories), 2)
	assert.NotEqual(t, acc1.GetId(), model.InvalidId)
	assert.NotEqual(t, acc2.GetId(), model.InvalidId)

	c.RemoveAccessory(acc2)
	assert.Equal(t, len(c.Accessories), 1)
}

func TestAccessoryId(t *testing.T) {
	acc1 := accessory.New(info)
	assert.Equal(t, acc1.GetId(), model.InvalidId)

	c := NewContainer()
	c.AddAccessory(acc1)
	id := acc1.GetId()
	assert.NotEqual(t, id, model.InvalidId)
	c.RemoveAccessory(acc1)
	c.AddAccessory(acc1)
	assert.Equal(t, acc1.GetId(), id)
}

func TestRemoveAccessory(t *testing.T) {
	accessory := accessory.New(info)
	c := NewContainer()
	c.AddAccessory(accessory)
	assert.Equal(t, len(c.Accessories), 1)
	c.RemoveAccessory(accessory)
	assert.Equal(t, len(c.Accessories), 0)
}
