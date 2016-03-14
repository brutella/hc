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

	a := accessory.New(info, accessory.TypeBridge)

	m := container.NewContainer()
	m.AddAccessory(a)

	aid := a.GetID()
	cid := a.Info.Name.GetID()
	values := idsString(aid, cid)
	controller := NewCharacteristicController(m)
	res, err := controller.HandleGetCharacteristics(values)

	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res)

	if err != nil {
		t.Fatal(err)
	}

	var chars data.Characteristics
	err = json.Unmarshal(b, &chars)

	if err != nil {
		t.Fatal(err)
	}

	if is, want := len(chars.Characteristics), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := chars.Characteristics[0].Value; x != "My Bridge" {
		t.Fatal(x)
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

	// find on characteristic with type TypePowerState
	var cid int64
	for _, s := range a.Accessory.Services {
		for _, c := range s.Characteristics {
			if c.Type == characteristic.TypePowerState {
				cid = c.ID
			}
		}
	}

	if cid == 0 {
		t.Fatal("characteristic not found")
	}

	char := data.Characteristic{AccessoryID: 1, CharacteristicID: cid, Value: true}
	var slice []data.Characteristic
	slice = append(slice, char)

	chars := data.Characteristics{Characteristics: slice}
	b, err := json.Marshal(chars)

	if err != nil {
		t.Fatal(err)
	}

	var buffer bytes.Buffer
	buffer.Write(b)

	controller := NewCharacteristicController(m)
	err = controller.HandleUpdateCharacteristics(&buffer, characteristic.TestConn)

	if err != nil {
		t.Fatal(err)
	}

	if is, want := a.IsOn(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
