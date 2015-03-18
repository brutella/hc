package netio

import (
	"github.com/brutella/hc/db"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	db, _ := db.NewDatabase(os.TempDir())
	client, err := NewClient("Test Client", db)
	assert.Nil(t, err)
	assert.True(t, len(client.PairPublicKey()) > 0)
	assert.True(t, len(client.PairPrivateKey()) > 0)

	entity := db.EntityWithName("Test Client")
	assert.Equal(t, entity.Name(), "Test Client")
	assert.True(t, len(entity.PublicKey()) > 0)
	assert.True(t, len(entity.PrivateKey()) > 0)
}
