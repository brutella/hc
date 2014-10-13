package container

import (
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"

    "testing"
    "github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
    info := model.Info{
        Name: "Accessory1",
        Serial: "001",
        Manufacturer: "Google",
        Model: "Accessory",
    }
    
    acc1 := accessory.New(info)
    
    info.Name = "Accessory2"
    acc2 := accessory.New(info)
    
    c := NewContainer()
    c.AddAccessory(acc1)
    c.AddAccessory(acc2)
    assert.Equal(t, len(c.Accessories), 2)
    
    assert.True(t, acc1.GetId() > 0)
    assert.True(t, acc2.GetId() > 0)
    
    c.RemoveAccessory(acc2)
    assert.Equal(t, len(c.Accessories), 1)
}