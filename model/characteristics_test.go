package model

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "encoding/json"
)

func TestStringCharacteristic(t *testing.T) {
    c := NewStringCharacteristic("my long string")
    result, _ := json.Marshal(c)
    assert.Equal(t, string(result), `{"iid":0,"type":"0","perms":["pr","pw"],"value":"my long string","format":"string"}`)
}

func TestBoolCharacteristic(t *testing.T) {
    c := NewBoolCharacteristic(true)
    result, _ := json.Marshal(c)
    assert.Equal(t, string(result), `{"iid":0,"type":"0","perms":["pr","pw"],"value":1,"format":"bool"}`)
}

func TestSerialNumberCharacteristic(t *testing.T) {
    serial := NewSerialNumberCharacteristic("ASDFG")
    result, _ := json.Marshal(serial)
    assert.Equal(t, string(result), `{"iid":0,"type":"30","perms":["pr"],"value":"ASDFG","format":"string"}`)
}

func TestManufacturerCharacteristic(t *testing.T) {
    manufacturer := NewManufacturerCharacteristic("Matthias Hochgatterer")
    result, _ := json.Marshal(manufacturer)
    assert.Equal(t, string(result), `{"iid":0,"type":"20","perms":["pr"],"value":"Matthias Hochgatterer","format":"string"}`)
}
 