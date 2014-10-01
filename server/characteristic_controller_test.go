package server

import (
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
    "github.com/brutella/hap/model/service"
    
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParseID(t *testing.T) {
    aid, cid, err := ParseAccessoryAndCharacterId("3.9")
    assert.Nil(t, err)
    assert.Equal(t, aid, 3)
    assert.Equal(t, cid, 9)
}

func TestGetCharacteristic(t *testing.T) {
    info_service := service.NewAccessoryInfoService( "My Bridge", "123-456-789", "Rev1", "Matthias H.")
    a := accessory.NewAccessory()
    a.AddService(info_service.Service)
    m := model.NewModel()
    m.AddAccessory(a)
    
    aid := a.Id
    cid := info_service.Name.Id
    
    controller := NewCharacteristicController(m)
    chars := controller.HandleGetCharacteristics(aid, cid)    
    for _, c := range chars.Characteristics {
        assert.Equal(t, c.Value, "My Bridge")
    }
}

func TestPutCharacteristic(t *testing.T) {
    info_service := service.NewAccessoryInfoService( "My Bridge", "123-456-789", "Rev1", "Matthias H.")
    a := accessory.NewAccessory()
    a.AddService(info_service.Service)
    m := model.NewModel()
    m.AddAccessory(a)
    
    aid := a.Id
    cid := info_service.Name.Id
    char := Characteristic{AccessoryId:aid, Id:cid, Value:"My"}
    slice := make([]Characteristic, 0)
    slice = append(slice, char)
    
    chars := Characteristics{Characteristics:slice}
    
    controller := NewCharacteristicController(m)
    err := controller.HandleUpdateCharacteristics(chars)
    assert.Nil(t, err)
    assert.Equal(t, info_service.Name.Value, "My")
}