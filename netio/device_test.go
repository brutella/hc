package netio

import (
	"github.com/brutella/hc/db"
	"os"
	"reflect"
	"testing"
)

func TestNewDevice(t *testing.T) {
	db, _ := db.NewDatabase(os.TempDir())
	client, err := NewDevice("Test Client", db)

	if err != nil {
		t.Fatal(err)
	}
	if x := len(client.PublicKey()); x == 0 {
		t.Fatal(x)
	}
	if x := len(client.PrivateKey()); x == 0 {
		t.Fatal(x)
	}

	entity := db.EntityWithName("Test Client")
	if is, want := entity.Name(), "Test Client"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if err != nil {
		t.Fatal(err)
	}
	if is, want := entity.PublicKey(), client.PublicKey(); reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := entity.PrivateKey(), client.PrivateKey(); reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
