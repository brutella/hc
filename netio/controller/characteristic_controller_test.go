package controller

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/container"
	_ "github.com/brutella/hc/model/service"
	"github.com/brutella/hc/netio/data"

	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGetCharacteristic(t *testing.T) {
	info := model.Info{
		Name:         "My Bridge",
		SerialNumber: "001",
		Manufacturer: "Google",
		Model:        "Bridge",
	}

	a := accessory.New(info)

	m := container.NewContainer()
	m.AddAccessory(a)

	aid := a.GetId()
	cid := a.Info.Name.GetId()
	values := getCharacteristicValues(aid, cid)
	controller := NewCharacteristicController(m)
	res, err := controller.HandleGetCharacteristics(values)
	assert.Nil(t, err)
	b, err := ioutil.ReadAll(res)
	assert.Nil(t, err)

	var chars data.Characteristics
	err = json.Unmarshal(b, &chars)
	assert.Nil(t, err)

	for _, c := range chars.Characteristics {
		assert.Equal(t, c.Value, "My Bridge")
	}
}

func TestPutCharacteristic(t *testing.T) {
	info := model.Info{
		Name:         "My Bridge",
		SerialNumber: "001",
		Manufacturer: "Google",
		Model:        "Bridge",
	}

	a := accessory.New(info)
	m := container.NewContainer()
	m.AddAccessory(a)

	aid := a.GetId()
	cid := a.Info.Name.GetId()
	char := data.Characteristic{AccessoryId: aid, Id: cid, Value: "My"}
	slice := make([]data.Characteristic, 0)
	slice = append(slice, char)

	chars := data.Characteristics{Characteristics: slice}
	b, err := json.Marshal(chars)
	assert.Nil(t, err)
	var buffer bytes.Buffer
	buffer.Write(b)

	controller := NewCharacteristicController(m)
	err = controller.HandleUpdateCharacteristics(&buffer)
	assert.Nil(t, err)
	assert.Equal(t, a.Info.Name.Value, "My")
}
