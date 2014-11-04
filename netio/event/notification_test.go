package event

import (
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
    
    "testing"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "bytes"
    "fmt"
)

func TestCharacteristicNotification(t *testing.T) {
    info := model.Info{
        Name: "My Bridge",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Bridge",
    }
    
    a := accessory.New(info)
        
    buffer, err :=  NotificationBody(a, a.Info.Name.Characteristic)
    assert.Nil(t, err)
    assert.NotNil(t, buffer)
    bytes, err := ioutil.ReadAll(buffer)
    assert.Nil(t, err)
    assert.Equal(t, string(bytes), `{"characteristics":[{"aid":0,"iid":6,"value":"My Bridge","ev":false}]}`)
}

func TestCharacteristicNotificationResponse(t *testing.T) {
    info := model.Info{
        Name: "My Bridge",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Bridge",
    }
    
    a := accessory.New(info)
        
    resp, err :=  NewNotification(a, a.Info.Name.Characteristic)
    assert.Nil(t, err)
    assert.NotNil(t, resp)
    
    var buffer = new(bytes.Buffer)
    resp.Write(buffer)
    
    bytes, err := ioutil.ReadAll(buffer)
    assert.Nil(t, err)
    fmt.Println(string(bytes))
}