package controller

import (
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
    "github.com/brutella/hap/model/service"
    "github.com/brutella/hap/netio/data"
    
    "testing"
    "github.com/stretchr/testify/assert"
    
    "net/url"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "bytes"
)

func TestGetCharacteristic(t *testing.T) {
    info_service := service.NewAccessoryInfo( "My Bridge", "123-456-789", "Rev1", "Matthias H.")
    a := accessory.NewAccessory()
    a.AddService(info_service.Service)
    m := model.NewModel()
    m.AddAccessory(a)
    
    aid := a.Id
    cid := info_service.Name.Id
    values := url.Values{}
    
    values.Set("id", fmt.Sprintf("%d.%d", aid, cid))
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
    info_service := service.NewAccessoryInfo( "My Bridge", "123-456-789", "Rev1", "Matthias H.")
    a := accessory.NewAccessory()
    a.AddService(info_service.Service)
    m := model.NewModel()
    m.AddAccessory(a)
    
    aid := a.Id
    cid := info_service.Name.Id
    char := data.Characteristic{AccessoryId:aid, Id:cid, Value:"My"}
    slice := make([]data.Characteristic, 0)
    slice = append(slice, char)
    
    chars := data.Characteristics{Characteristics:slice}
    b, err := json.Marshal(chars)
    assert.Nil(t, err)
    var buffer bytes.Buffer
    buffer.Write(b)
    
    controller := NewCharacteristicController(m)
    err = controller.HandleUpdateCharacteristics(&buffer)
    assert.Nil(t, err)
    assert.Equal(t, info_service.Name.Value, "My")
}