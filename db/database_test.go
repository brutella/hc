package db

import (    
	"testing"
    "github.com/stretchr/testify/assert"
    "os"
)

func TestLoadUndefinedClient(t *testing.T) {
    db, _ := NewDatabase(os.TempDir())
    client := db.ClientWithName("My Name")
    assert.Nil(t, client)
}

func TestLoadClient(t *testing.T) {
    db, _ := NewDatabase(os.TempDir())
    
    db.SaveClient(NewClient("My Name", []byte{0x01}))
    client := db.ClientWithName("My Name")
    assert.NotNil(t, client)
    assert.Equal(t, client.PublicKey(), []byte{0x01})
}

func TestDeleteClient(t *testing.T) {
    db, _ := NewDatabase(os.TempDir())
    
    c := NewClient("My Name", []byte{0x01})
    db.SaveClient(c)
    db.DeleteClient(c)
    client := db.ClientWithName("My Name")
    assert.Nil(t, client)
}