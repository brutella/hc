package tlv8

import (
	"testing"
)

type person struct {
	Name    string  `tlv8:"1"`
	Aliases []alias `tlv8:"-"`
}

type alias struct {
	Name string `tlv8:"3"`
}

var p = person{"Matthias", []alias{alias{"brutella"}, alias{"adsf"}}}

func TestUnmarshalPerson(t *testing.T) {
	tlv8, err := Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	var pers person
	err = Unmarshal(tlv8, &pers)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := pers.Name, "Matthias"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := len(pers.Aliases), 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := pers.Aliases[0].Name, "brutella"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := pers.Aliases[1].Name, "adsf"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
