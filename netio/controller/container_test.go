package controller

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/container"

	"bytes"
	"encoding/json"
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

	a := accessory.New(info, accessory.TypeOther)

	m := container.NewContainer()
	m.AddAccessory(a)

	controller := NewContainerController(m)

	var b bytes.Buffer
	r, err := controller.HandleGetAccessories(&b)

	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("no response")
	}

	bytes, _ := ioutil.ReadAll(r)
	var returnedContainer container.Container
	if err := json.Unmarshal(bytes, &returnedContainer); err != nil {
		t.Fatal(err)
	}

	if returnedContainer.Equal(m) == false {
		t.Fatal("containers not the same")
	}
}
