package db

import (
    "github.com/brutella/hap"
    
	"testing"
    "github.com/stretchr/testify/assert"
    "os"
)

func TestSaveAndLoadParty(t *testing.T) {
    storage, _ := hap.NewFileStorage(os.TempDir())
    db := NewManager(storage)
    
    party := NewParty("My Name", "My Serial", []byte{0x01}, []byte{0x02})
    db.SaveParty(party)
    loaded_party := db.PartyWithName("My Name")
    assert.NotNil(t, loaded_party)
    assert.Equal(t, loaded_party.SerialNumber, "My Serial")
    assert.Equal(t, loaded_party.PublicKey, []byte{0x01})
    assert.Equal(t, loaded_party.SecretKey, []byte{0x02})
}

func TestSaveAndLoadRemoteParty(t *testing.T) {
    storage, _ := hap.NewFileStorage(os.TempDir())
    db := NewManager(storage)
    
    party := NewParty("My Name", "", []byte{0x01}, nil)
    db.SaveParty(party)
    loaded_party := db.PartyWithName("My Name")
    assert.NotNil(t, loaded_party)
    assert.Equal(t, loaded_party.SerialNumber, "")
    assert.Equal(t, loaded_party.PublicKey, []byte{0x01})
    assert.Nil(t, loaded_party.SecretKey)
}
