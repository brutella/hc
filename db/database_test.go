package db

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadUndefinedEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	entity := db.EntityWithName("My Name")
	assert.Nil(t, entity)
}

func TestLoadEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	db.SaveEntity(NewEntity("My Name", []byte{0x01}, []byte{0x02}))
	entity := db.EntityWithName("My Name")
	assert.NotNil(t, entity)
	assert.Equal(t, entity.PublicKey(), []byte{0x01})
	assert.Equal(t, entity.PrivateKey(), []byte{0x02})
}

func TestLoadEntityWithPublicKeyOnly(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	db.SaveEntity(NewEntity("Entity", []byte{0x03}, nil))
	entity := db.EntityWithName("Entity")
	assert.NotNil(t, entity)
	assert.Equal(t, entity.PublicKey(), []byte{0x03}, nil)
	assert.Nil(t, entity.PrivateKey())
}

func TestDeleteEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	c := NewEntity("My Name", []byte{0x01}, nil)
	db.SaveEntity(c)
	db.DeleteEntity(c)
	entity := db.EntityWithName("My Name")
	assert.Nil(t, entity)
}

func TestLoadDns(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	dns := NewDns("My Name", 10, 20)
	db.SaveDns(dns)
	entity := db.DnsWithName("My Name")
	assert.NotNil(t, entity)
	assert.Equal(t, entity.Configuration(), 10)
	assert.Equal(t, entity.State(), 20)
}

func TestDeleteDns(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	dns := NewDns("My Name", 10, 20)
	db.SaveDns(dns)
	db.DeleteDns(dns)
	assert.Nil(t, db.DnsWithName("My Name"))
}
