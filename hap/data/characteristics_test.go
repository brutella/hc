package data

import (
	"github.com/xiam/to"

	"encoding/json"
	"reflect"
	"testing"
)

func TestEventCharacteristicFromJSON(t *testing.T) {
	var b = []byte(`{"characteristics":[{"aid":2,"iid":13,"status":1,"ev":true}]}`)
	var cs Characteristics

	if err := json.Unmarshal(b, &cs); err != nil {
		t.Fatal(err)
	}
	if x := len(cs.Characteristics); x != 1 {
		t.Fatal(x)
	}

	c := cs.Characteristics[0]

	if x := c.AccessoryID; x != 2 {
		t.Fatal(x)
	}
	if x := c.CharacteristicID; x != 13 {
		t.Fatal(x)
	}

	if x, ok := c.Events.(bool); ok {
		if !x {
			t.Fatalf("want=true is=%v", x)
		}
	} else {
		t.Fatalf("invalid events type %v", reflect.TypeOf(x))
	}

	if x := to.Int64(c.Status); x != 1 {
		t.Fatal(x)
	}

	if c.Value != nil {
		t.Fatalf("want=nil is=%v", c.Value)
	}
}
