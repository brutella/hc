package pair

import (
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/util"

	"os"
	"testing"
)

func TestUnknownPairingMethod(t *testing.T) {
	tlv8 := util.NewTLV8Container()
	tlv8.SetByte(TagPairingMethod, 0x09)

	database, _ := db.NewDatabase(os.TempDir())
	controller := NewPairingController(database)

	out, err := controller.Handle(tlv8)

	if err == nil {
		t.Fatal("expected error for unknown pairing method")
	}
	if out != nil {
		t.Fatal(out)
	}
}

func TestAddPairing(t *testing.T) {
	in := util.NewTLV8Container()
	in.SetByte(TagPairingMethod, PairingMethodAdd.Byte())
	in.SetByte(TagSequence, 0x01)
	in.SetString(TagUsername, "Unit Test")
	in.SetBytes(TagPublicKey, []byte{0x01, 0x02})

	database, _ := db.NewDatabase(os.TempDir())
	controller := NewPairingController(database)

	out, err := controller.Handle(in)
	if err != nil {
		t.Fatal(err)
	}
	if out == nil {
		t.Fatal("no response")
	}
	if is, want := out.GetByte(TagSequence), byte(0x2); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestDeletePairing(t *testing.T) {
	username := "Unit Test"
	entity := db.NewEntity(username, []byte{0x01, 0x02}, nil)
	database, _ := db.NewDatabase(os.TempDir())
	database.SaveEntity(entity)

	in := util.NewTLV8Container()
	in.SetByte(TagPairingMethod, PairingMethodDelete.Byte())
	in.SetByte(TagSequence, 0x01)
	in.SetString(TagUsername, username)

	controller := NewPairingController(database)

	out, err := controller.Handle(in)
	if err != nil {
		t.Fatal(err)
	}
	if out == nil {
		t.Fatal("no response")
	}
	if is, want := out.GetByte(TagSequence), byte(0x2); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	savedEntity := database.EntityWithName(username)
	if savedEntity != nil {
		t.Fatal(savedEntity)
	}
}
