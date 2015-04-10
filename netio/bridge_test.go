package netio

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBridge(t *testing.T) {
	storage, _ := common.NewTempFileStorage()
	info := NewAccessoryInfo("Test Bridge", "719-47-107", "Matthias H.", storage)
	db := db.NewDatabaseWithStorage(storage)
	bridge, err := NewBridge(info, db)
	assert.Nil(t, err)
	assert.True(t, len(bridge.PairPublicKey()) > 0)
	assert.True(t, len(bridge.PairPrivateKey()) > 0)
}
