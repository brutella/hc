package common

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
    dir, _ := filepath.Abs(filepath.Join(os.TempDir(), "hap"))
    storage, err := NewFileStorage(dir)
    assert.Nil(t, err)
    assert.NotNil(t, storage)
        
    err = storage.Set("test", []byte("ASDF"))
    assert.Nil(t, err)
    
    path := filepath.Join(dir, "test")
    f, err := os.OpenFile(path, os.O_RDONLY, 0776)
    assert.Nil(t, err)
    assert.NotNil(t, f)
    
    defer f.Close()
    
    var buffer []byte = make([]byte, 32)
    n, _ := f.Read(buffer)
    assert.Equal(t, buffer[:n], []byte("ASDF"))
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