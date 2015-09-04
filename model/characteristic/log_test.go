package characteristic

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	// "reflect"
)

func TestLogs(t *testing.T) {
	dir := os.TempDir()
	file := filepath.Join(dir, "test.log")
	f, err := os.Create(file)

	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString("This is a test string"); err != nil {
		t.Fatal(err)
	}

	n := NewLog(file)

	if is, want := n.Type, TypeLogs; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	b, err := json.Marshal(n)

	if err != nil {
		t.Fatal(err)
	}

	h := map[string]interface{}{}

	if err := json.Unmarshal(b, &h); err != nil {
		t.Fatal(err)
	}

	s, ok := h["value"].(string)

	if ok == false {
		t.Fatal(ok)
	}

	b, err = bytesFromTLV8Base64(s)

	if err != nil {
		t.Fatal(err)
	}
	if is, want := string(b), "This is a test string"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
