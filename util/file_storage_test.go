package util

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFileStorage(t *testing.T) {
	storage, err := NewTempFileStorage()
	if err != nil {
		t.Fatal(err)
	}

	if err := storage.Set("test", []byte("ASDF")); err != nil {
		t.Fatal(err)
	}

	read, err := storage.Get("test")

	if err != nil {
		t.Fatal(err)
	}
	if is, want := read, []byte("ASDF"); reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if err := storage.Delete("test"); err != nil {
		t.Fatal(err)
	}
}

func TestStoreInSubdirectory(t *testing.T) {
	dir, _ := filepath.Abs(filepath.Join(os.TempDir(), "hap"))
	storage, err := NewFileStorage(dir)

	if err != nil {
		t.Fatal(err)
	}

	err = storage.Set("test", []byte("ASDF"))

	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(dir, "test")
	f, err := os.OpenFile(path, os.O_RDONLY, 0776)

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	var buffer = make([]byte, 32)
	n, _ := f.Read(buffer)

	if is, want := buffer[:n], []byte("ASDF"); reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestDeleteUndefined(t *testing.T) {
	storage, err := NewTempFileStorage()

	if err != nil {
		t.Fatal(err)
	}
	if err := storage.Delete("test"); err == nil {
		t.Fatal("expected error")
	}
}

func TestGetUndefined(t *testing.T) {
	storage, err := NewTempFileStorage()

	if err != nil {
		t.Fatal(err)
	}
	if _, err := storage.Get("test"); err == nil {
		t.Fatal("expected error")
	}
}
