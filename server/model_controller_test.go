package server

import (
	"testing"
    "github.com/brutella/hap/model"
    "github.com/stretchr/testify/assert"
    "bytes"
    "io/ioutil"
    "encoding/json"
)

func TestGetAccessories(t *testing.T) {
    info_service := model.NewAccessoryInfoService("123-456-789", "Rev1", "Matthias H.", "My Bridge")
    a := model.NewAccessory()
    a.AddService(info_service.Service)
    m := model.NewModel()
    m.AddAccessory(a)
    
    controller := NewModelController(m)
    
    var b bytes.Buffer
    r, err := controller.HandleGetAccessories(&b)
    assert.Nil(t, err)
    assert.NotNil(t, r)
    
    bytes, _ := ioutil.ReadAll(r)
    var returnedModel model.Model
    err = json.Unmarshal(bytes, &returnedModel)
    assert.Nil(t, err)
    assert.True(t, returnedModel.Equal(m))
}