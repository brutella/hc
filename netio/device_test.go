package netio

import (
	"github.com/brutella/hc/db"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewDevice(t *testing.T) {
	db, _ := db.NewDatabase(os.TempDir())
	client, err := NewDevice("Test Client", db)
	assert.Nil(t, err)
	assert.True(t, len(client.PublicKey()) > 0)
	assert.True(t, len(client.PrivateKey()) > 0)

	entity := db.EntityWithName("Test Client")
	assert.Equal(t, entity.Name(), "Test Client")
	assert.True(t, len(entity.PublicKey()) > 0)
	assert.True(t, len(entity.PrivateKey()) > 0)
}
