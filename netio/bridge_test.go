package netio

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewBridge(t *testing.T) {
	storage, _ := common.NewFileStorage(os.TempDir())
	info := NewBridgeInfo("Test Bridge", "719-47-107", "Matthias H.", storage)
	db := db.NewDatabaseWithStorage(storage)
	bridge, err := NewBridge(info, db)
	assert.Nil(t, err)
	assert.True(t, len(bridge.PublicKey()) > 0)
	assert.True(t, len(bridge.PrivateKey()) > 0)
}
