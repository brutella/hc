package server

import (
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
    "github.com/brutella/hap/model/service"
    
	"testing"    
    "github.com/stretchr/testify/assert"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "fmt"
)

func TestGetAccessories(t *testing.T) {
    info_service := service.NewAccessoryInfoService("123-456-789", "Rev1", "Matthias H.", "My Bridge")
    a := accessory.NewAccessory()
    a.AddService(info_service.Service)
    m := model.NewModel()
    m.AddAccessory(a)
    
    controller := NewModelController(m)
    
    var b bytes.Buffer
    r, err := controller.HandleGetAccessories(&b)
    assert.Nil(t, err)
    assert.NotNil(t, r)
    
    bytes, _ := ioutil.ReadAll(r)
    fmt.Println(string(bytes))
    var returnedModel model.Model
    err = json.Unmarshal(bytes, &returnedModel)
    assert.Nil(t, err)
    assert.True(t, returnedModel.Equal(m))
}