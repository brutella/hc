package netio

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunkedWriter(t *testing.T) {
	var b bytes.Buffer
	wr := NewChunkedWriter(&b, 2)
	n, err := wr.Write([]byte("Hello World"))
	assert.Nil(t, err)
	assert.Equal(t, n, 11)
	assert.Equal(t, string(b.Bytes()), "Hello World")
}
