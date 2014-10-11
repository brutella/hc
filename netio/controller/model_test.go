package controller

import (
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/container"
    "github.com/brutella/hap/model/accessory"
    _"github.com/brutella/hap/model/service"
    
	"testing"    
    "github.com/stretchr/testify/assert"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "fmt"
)

func TestGetAccessories(t *testing.T) {
    info := model.Info{
            Name: "My Accessory",
            Serial: "001",
            Manufacturer: "Google",
            Model: "Accessory",
        }
    
    a := accessory.New(info)
        
    m := container.NewContainer()
    m.AddAccessory(a)
    
    controller := NewContainerController(m)
    
    var b bytes.Buffer
    r, err := controller.HandleGetAccessories(&b)
    assert.Nil(t, err)
    assert.NotNil(t, r)
    
    bytes, _ := ioutil.ReadAll(r)
    fmt.Println(string(bytes))
    var returnedContainer container.Container
    err = json.Unmarshal(bytes, &returnedContainer)    
    assert.Nil(t, err)
    assert.True(t, returnedContainer.Equal(m))
}