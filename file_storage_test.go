package hap

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "os"
    "path/filepath"
)

func TestFileStorage(t *testing.T) {
    storage, err := NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    assert.NotNil(t, storage)
    
    err = storage.Set("test", []byte("ASDF"))
    assert.Nil(t, err)
    
    read, err := storage.Get("test")
    assert.Nil(t, err)
    assert.Equal(t, read, []byte("ASDF"))
    
    assert.Nil(t, storage.Delete("test"))
}

func TestStoreInSubdirectory (t *testing.T){
    storage, err := NewFileStorage(filepath.Join(os.TempDir(), "hap"))
    assert.Nil(t, err)
    assert.NotNil(t, storage)
    
    err = storage.Set("test", []byte("ASDF"))
    assert.Nil(t, err)
    
    read, err := storage.Get("test")
    assert.Nil(t, err)
    assert.Equal(t, read, []byte("ASDF"))
    
    assert.Nil(t, storage.Delete("test"))
}

func TestDeleteUndefined(t *testing.T) {
    storage, err := NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    assert.NotNil(t, storage.Delete("test"))
}

func TestGetUndefined(t *testing.T) {
    storage, err := NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    
    _, err = storage.Get("test")
    assert.NotNil(t, err)
}