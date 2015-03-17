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
	assert.True(t, len(client.PublicKey()) > 0)
	assert.True(t, len(client.PrivateKey()) > 0)
}
