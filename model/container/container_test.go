package container

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"

	"github.com/stretchr/testify/assert"
	"testing"
)

var info = model.Info{
	Name:         "Accessory1",
	SerialNumber: "001",
	Manufacturer: "Google",
	Model:        "Accessory",
}

func TestContainer(t *testing.T) {
	acc1 := accessory.New(info)
	assert.Equal(t, acc1.GetID(), model.InvalidID)

	info.Name = "Accessory2"
	acc2 := accessory.New(info)
	assert.Equal(t, acc2.GetID(), model.InvalidID)

	c := NewContainer()
	c.AddAccessory(acc1)
	c.AddAccessory(acc2)

	assert.Equal(t, len(c.Accessories), 2)
	assert.NotEqual(t, acc1.GetID(), model.InvalidID)
	assert.NotEqual(t, acc2.GetID(), model.InvalidID)
	assert.NotEqual(t, acc1.GetID(), acc2.GetID())

	c.RemoveAccessory(acc2)
	assert.Equal(t, len(c.Accessories), 1)
}

func TestValidAccessoryID(t *testing.T) {
	acc1 := accessory.New(info)
	assert.Equal(t, acc1.GetID(), model.InvalidID)

	c := NewContainer()
	c.AddAccessory(acc1)
	id := acc1.GetID()
	assert.NotEqual(t, id, model.InvalidID)
	c.RemoveAccessory(acc1)
	c.AddAccessory(acc1)
	assert.Equal(t, acc1.GetID(), id)
}

func TestRemoveAccessory(t *testing.T) {
	accessory := accessory.New(info)
	c := NewContainer()
	c.AddAccessory(accessory)
	assert.Equal(t, len(c.Accessories), 1)
	c.RemoveAccessory(accessory)
	assert.Equal(t, len(c.Accessories), 0)
}
