package data

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventCharacteristicFromJSON(t *testing.T) {
	var b = []byte(`{"characteristics":[{"aid":2,"iid":13,"ev":true}]}`)
	var chars Characteristics
	err := json.Unmarshal(b, &chars)
	assert.Nil(t, err)
	assert.Equal(t, len(chars.Characteristics), 1)
	var char = chars.Characteristics[0]
	assert.Equal(t, char.AccessoryID, 2)
	assert.Equal(t, char.ID, 13)
	assert.Equal(t, char.Events, true)
	assert.Nil(t, char.Value)
}
