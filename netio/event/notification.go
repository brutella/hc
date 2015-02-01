package event

import (
	"bytes"
	"encoding/json"
	"github.com/brutella/hap/model/accessory"
	"github.com/brutella/hap/model/characteristic"
	"github.com/brutella/hap/netio"
	"github.com/brutella/hap/netio/data"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewNotification(a *accessory.Accessory, c *characteristic.Characteristic) (*http.Response, error) {
	body, err := NotificationBody(a, c)
	if err != nil {
		return nil, err
	}

	resp := new(http.Response)
	resp.Status = "200 OK"
	resp.StatusCode = http.StatusOK
	resp.ProtoMajor = 1
	resp.ProtoMinor = 0
	resp.Body = ioutil.NopCloser(body)
	resp.ContentLength = int64(body.Len())
	resp.Header = map[string][]string{}
	resp.Header.Set("Content-Type", netio.HTTPContentTypeHAPJson)
	// (brutella) Not sure if Date header must be set
	// resp.Header.Set("Date", netio.CurrentRFC1123Date())
	resp.Proto = "EVENT/1.0"

	return resp, nil
}

func FixProtocolSpecifier(b []byte) []byte {
	return []byte(strings.Replace(string(b), "HTTP/1.0", "EVENT/1.0", 1))
}

func NotificationBody(a *accessory.Accessory, c *characteristic.Characteristic) (*bytes.Buffer, error) {
	chars := data.NewCharacteristics()
	char := data.Characteristic{AccessoryId: a.GetId(), Id: c.GetId(), Value: c.GetValue()}
	chars.AddCharacteristic(char)

	result, err := json.Marshal(chars)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	b.Write(result)
	return &b, err
}
