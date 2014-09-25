package hap

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "encoding/json"
    "fmt"
)

func TestCharacteristicJSON(t *testing.T) {
    c := NewCharacteristic()
    result, err := json.Marshal(c)
    assert.Nil(t, err)
    assert.NotNil(t, result)
    fmt.Println(string(result))
}
 