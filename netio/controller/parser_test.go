package controller

import (
	"testing"
)

func TestParseID(t *testing.T) {
	aid, cid, err := ParseAccessoryAndCharacterID("3.9")
	if err != nil {
		t.Fatal(err)
	}
	if aid != 3 {
		t.Fatal(aid)
	}
	if cid != 9 {
		t.Fatal(cid)
	}
}

func TestParseInvalidID(t *testing.T) {
	_, _, err := ParseAccessoryAndCharacterID("random")
	if err == nil {
		t.Fatal("err nil")
	}
}
