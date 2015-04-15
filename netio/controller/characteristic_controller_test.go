package controller

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/container"
	"github.com/brutella/hc/model/service"
	"github.com/brutella/hc/netio/data"

	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/url"
	"testing"
)

func idsString(accessoryID, characteristicID int64) url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d.%d", accessoryID, characteristicID))

	return values
}

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

	aid := a.GetID()
	cid := a.Info.Name.GetID()
	values := idsString(aid, cid)
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

func toSwitchService(obj interface{}) *service.Switch {
	return obj.(*service.Switch)
}

func TestPutCharacteristic(t *testing.T) {
	info := model.Info{
		Name:         "My Switch",
		SerialNumber: "001",
		Manufacturer: "Google",
		Model:        "Bridge",
	}

	a := accessory.NewSwitch(info)
	a.SetOn(false)

	m := container.NewContainer()
	m.AddAccessory(a.Accessory)

	// find on characteristic with type CharTypePowerState
	var cid int64
	for _, s := range a.Accessory.Services {
		for _, c := range s.Characteristics {
			if c.Type == characteristic.CharTypePowerState {
				cid = c.ID
			}
		}
	}
	assert.NotEqual(t, 0, cid, "Could not find power state characteristic")
	char := data.Characteristic{AccessoryID: 1, ID: cid, Value: true}
	var slice []data.Characteristic
	slice = append(slice, char)

	chars := data.Characteristics{Characteristics: slice}
	b, err := json.Marshal(chars)
	assert.Nil(t, err)
	var buffer bytes.Buffer
	buffer.Write(b)

	controller := NewCharacteristicController(m)
	err = controller.HandleUpdateCharacteristics(&buffer)
	assert.Nil(t, err)
	assert.Equal(t, a.IsOn(), true)
}
