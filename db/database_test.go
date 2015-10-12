package db

import (
	"reflect"
	"testing"
)

func TestLoadUndefinedEntity(t *testing.T) {
	db, _ := NewTempDatabase()

	if _, err := db.EntityWithName("My Name"); err == nil {
		t.Fatal("expected error")
	}
}

func TestGetEntity(t *testing.T) {
	db, _ := NewTempDatabase()
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

func TestGetEntityWithPublicKeyOnly(t *testing.T) {
	db, _ := NewTempDatabase()
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
	db, _ := NewTempDatabase()
	c := NewEntity("My Name", []byte{0x01}, nil)
	db.SaveEntity(c)
	db.DeleteEntity(c)
	if _, err := db.EntityWithName("My Name"); err == nil {
		t.Fatal("expected error")
	}
}

func TestGetEntities(t *testing.T) {
	db, _ := NewTempDatabase()
	e1 := NewEntity("Entity 1", []byte{0x01}, []byte{0x02})
	e2 := NewEntity("Entity 2", []byte{0x01}, []byte{0x02})

	db.SaveEntity(e1)
	db.SaveEntity(e2)

	var es []Entity
	var err error

	if es, err = db.Entities(); err != nil {
		t.Fatal(err)
	}

	if x := es; reflect.DeepEqual(x, []Entity{e1, e2}) == false {
		t.Fatal(x)
	}
}
