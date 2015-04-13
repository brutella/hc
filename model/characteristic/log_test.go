package characteristic

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLogs(t *testing.T) {
	dir := os.TempDir()
	file := filepath.Join(dir, "test.log")
	f, err := os.Create(file)
	assert.Nil(t, err)
	_, err = f.WriteString("This is a test string")
	assert.Nil(t, err)

	n := NewLog(file)
	assert.Equal(t, n.Type, CharTypeLogs)
	_, err = json.Marshal(n)
	assert.Nil(t, err)
}
