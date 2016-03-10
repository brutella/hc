package netio

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/container"

	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

var info = model.Info{
	Name:         "My Bridge",
	SerialNumber: "001",
	Manufacturer: "Google",
	Model:        "Bridge",
}

func TestCharacteristicNotification(t *testing.T) {
	a := accessory.New(info, accessory.TypeOther)
	c := container.NewContainer()
	c.AddAccessory(a)

	buffer, err := Body(a, a.Info.Name.Characteristic)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(buffer)

	if err != nil {
		t.Fatal(err)
	}
	if is, want := string(bytes), `{"characteristics":[{"aid":1,"iid":2,"value":"My Bridge"}]}`; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCharacteristicNotificationResponse(t *testing.T) {
	a := accessory.New(info, accessory.TypeOther)
	resp, err := New(a, a.Info.Name.Characteristic)

	if err != nil {
		t.Fatal(err)
	}

	var buffer = new(bytes.Buffer)
	resp.Write(buffer)

	bytes, err := ioutil.ReadAll(buffer)

	if err != nil {
		t.Fatal(err)
	}
	bytes = FixProtocolSpecifier(bytes)
	if x := string(bytes); strings.HasPrefix(x, "EVENT/1.0 200 OK") == false {
		t.Fatal(x)
	}
}
