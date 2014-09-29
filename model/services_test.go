package hk

// import (
//     "testing"
//     "github.com/stretchr/testify/assert"
//     "encoding/json"
//     "fmt"
// )
//
// func TestAccessoryInfoService(t *testing.T) {
//     // serialNumber, modelName, manufacturerName, accessoryName string
//     service := NewAccessoryInfoService("123-456-789", "Rev1", "Matthias H.", "My Accessory")
//     result, _ := json.Marshal(service)
//     assert.NotNil(t, result)
//     assert.Equal(t, string(result), `{"iid":0,"type":"3E","characteristics":[{"iid":0,"type":"14","perms":["pw"],"value":0,"format":"bool"},{"iid":0,"type":"30","perms":["pr"],"value":"123-456-789","format":"string"},{"iid":0,"type":"21","perms":["pr"],"value":"Rev1","format":"string"},{"iid":0,"type":"20","perms":["pr"],"value":"Matthias H.","format":"string"},{"iid":0,"type":"23","perms":["pr"],"value":"My Accessory","format":"string"}]}`)
// }
//
// func TestThermostatService(t *testing.T) {
//     // serialNumber, modelName, manufacturerName, accessoryName string
//     service := NewThermostatService("Living Room", 25, 0, 100, 1.0)
//     result, _ := json.Marshal(service)
//     assert.NotNil(t, result)
//     fmt.Println(string(result))
// }