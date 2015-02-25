package controller

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/container"

	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGetAccessories(t *testing.T) {
	info := model.Info{
		Name:         "My Accessory",
		SerialNumber: "001",
		Manufacturer: "Google",
		Model:        "Accessory",
	}

	a := accessory.New(info)

	m := container.NewContainer()
	m.AddAccessory(a)

	controller := NewContainerController(m)

	var b bytes.Buffer
	r, err := controller.HandleGetAccessories(&b)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	bytes, _ := ioutil.ReadAll(r)
	var returnedContainer container.Container
	err = json.Unmarshal(bytes, &returnedContainer)
	assert.Nil(t, err)
	assert.True(t, returnedContainer.Equal(m))
}
