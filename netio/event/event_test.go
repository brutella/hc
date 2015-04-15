package event

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"github.com/brutella/hc/model/container"

	"bytes"
	"github.com/stretchr/testify/assert"
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
	a := accessory.New(info)
	c := container.NewContainer()
	c.AddAccessory(a)

	buffer, err := Body(a, a.Info.Name.Characteristic)
	assert.Nil(t, err)
	assert.NotNil(t, buffer)
	bytes, err := ioutil.ReadAll(buffer)
	assert.Nil(t, err)
	assert.Equal(t, string(bytes), `{"characteristics":[{"aid":1,"iid":2,"value":"My Bridge"}]}`)
}

func TestCharacteristicNotificationResponse(t *testing.T) {
	a := accessory.New(info)
	resp, err := New(a, a.Info.Name.Characteristic)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	var buffer = new(bytes.Buffer)
	resp.Write(buffer)

	bytes, err := ioutil.ReadAll(buffer)
	assert.Nil(t, err)
	bytes = FixProtocolSpecifier(bytes)
	str := string(bytes)
	assert.True(t, strings.HasPrefix(str, "EVENT/1.0 200 OK"))
}
