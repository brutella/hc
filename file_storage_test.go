package gohap

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "os"
)

func TestStoreInFile(t *testing.T) {
    storage := NewFileStorage(os.TempDir())
    assert.NotNil(t, storage)
    
    err := storage.Update("test_key", []byte("ASDF"))
    assert.Nil(t, err)
    
    read, err := storage.Get("test_key")
    assert.Nil(t, err)
    assert.Equal(t, read, []byte("ASDF"))
    
    assert.Nil(t, storage.Delete("test_key"))
}