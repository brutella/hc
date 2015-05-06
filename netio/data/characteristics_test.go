package data

import (
	"encoding/json"
	"testing"
)

func TestEventCharacteristicFromJSON(t *testing.T) {
	var b = []byte(`{"characteristics":[{"aid":2,"iid":13,"ev":true}]}`)
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
	if x := c.ID; x != 13 {
		t.Fatal(x)
	}
	if x, ok := c.Events.(bool); ok {
		if !x {
			t.Fatalf("want=true is=%v", x)
		}
	} else {
		t.Fatal("invalid Events type")
	}
	if c.Value != nil {
		t.Fatalf("want=nil is=%v", c.Value)
	}
}
