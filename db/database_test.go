package db

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadUndefinedEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())

	if _, err := db.EntityWithName("My Name"); err == nil {
		t.Fatal("expected error")
	}
}

func TestLoadEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	db.SaveEntity(NewEntity("My Name", []byte{0x01}, []byte{0x02}))
	var e Entity
	var err error

	if e, err = db.EntityWithName("My Name"); err != nil {
		t.Fatal(err)
	}

	if x := e.PublicKey; reflect.DeepEqual(x, []byte{0x01}) == false {
		t.Fatal(x)
	}

	if x := e.PrivateKey; reflect.DeepEqual(x, []byte{0x02}) == false {
		t.Fatal(x)
	}
}

func TestLoadEntityWithPublicKeyOnly(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	db.SaveEntity(NewEntity("Entity", []byte{0x03}, nil))

	var e Entity
	var err error

	if e, err = db.EntityWithName("Entity"); err != nil {
		t.Fatal(err)
	}

	if x := e.PublicKey; reflect.DeepEqual(x, []byte{0x03}) == false {
		t.Fatal(x)
	}

	if x := e.PrivateKey; x != nil {
		t.Fatal(x)
	}
}

func TestDeleteEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	c := NewEntity("My Name", []byte{0x01}, nil)
	db.SaveEntity(c)
	db.DeleteEntity(c)
	if _, err := db.EntityWithName("My Name"); err == nil {
		t.Fatal("expected error")
	}
}
