package service

import (
	"encoding/json"
	"testing"
)

func TestMinimalService(t *testing.T) {
	s := New(TypeOutlet)

	if buf, err := json.Marshal(s); err != nil {
		t.Fatal(err)
	} else {
		if is, want := string(buf), "{\"iid\":0,\"type\":\"47\",\"characteristics\":[]}"; is != want {
			t.Fatalf("%v != %v", is, want)
		}
	}
}

func TestPrimaryService(t *testing.T) {
	s := New(TypeOutlet)
	s.SetPrimary(true)

	if buf, err := json.Marshal(s); err != nil {
		t.Fatal(err)
	} else {
		if is, want := string(buf), "{\"iid\":0,\"type\":\"47\",\"characteristics\":[],\"primary\":true}"; is != want {
			t.Fatalf("%v != %v", is, want)
		}
	}
}

func TestHiddenService(t *testing.T) {
	s := New(TypeOutlet)
	s.SetHidden(true)

	if buf, err := json.Marshal(s); err != nil {
		t.Fatal(err)
	} else {
		if is, want := string(buf), "{\"iid\":0,\"type\":\"47\",\"characteristics\":[],\"hidden\":true}"; is != want {
			t.Fatalf("%v != %v", is, want)
		}
	}
}

func TestLinkedService(t *testing.T) {
	s := New(TypeOutlet)
	s.ID = 1
	fan := New(TypeFan)
	fan.ID = 2

	s.AddLinkedService(fan)

	if buf, err := json.Marshal(s); err != nil {
		t.Fatal(err)
	} else {
		if is, want := string(buf), "{\"iid\":1,\"type\":\"47\",\"characteristics\":[],\"linked\":[2]}"; is != want {
			t.Fatalf("%v != %v", is, want)
		}
	}
}
