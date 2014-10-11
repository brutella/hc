package controller

import (
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/model"
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
    
    jerr := err.(*json.UnmarshalTypeError)
    fmt.Printf("Unexpected value: %s\n", (*jerr).Value)

    // Expected type: uint
    fmt.Printf("Unexpected type: %v\n", (*jerr).Type)
    
    assert.Nil(t, err)
    assert.True(t, returnedModel.Equal(m))
}