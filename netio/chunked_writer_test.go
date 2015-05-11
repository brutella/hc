package netio

import (
	"bytes"
	"testing"
)

func TestChunkedWriter(t *testing.T) {
	var b bytes.Buffer
	wr := NewChunkedWriter(&b, 2)
	n, err := wr.Write([]byte("Hello World"))
	if err != nil {
		t.Fatal(err)
	}
	if is, want := n, 11; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := string(b.Bytes()), "Hello World"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
