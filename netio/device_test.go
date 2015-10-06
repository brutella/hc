package netio

import (
	"github.com/brutella/hc/db"
	"os"
	"reflect"
	"testing"
)

func TestNewDevice(t *testing.T) {
	var e db.Entity
	var client Device
	var err error

	database, _ := db.NewDatabase(os.TempDir())
	client, err = NewDevice("Test Client", database)

	if err != nil {
		t.Fatal(err)
	}
	if x := len(client.PublicKey()); x == 0 {
		t.Fatal(x)
	}
	if x := len(client.PrivateKey()); x == 0 {
		t.Fatal(x)
	}

	if e, err = database.EntityWithName("Test Client"); err != nil {
		t.Fatal(err)
	}

	if is, want := e.Name, "Test Client"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if err != nil {
		t.Fatal(err)
	}
	if is, want := e.PublicKey, client.PublicKey(); reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := e.PrivateKey, client.PrivateKey(); reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
