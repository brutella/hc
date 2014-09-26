package crypto

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "bytes"
    "io/ioutil"
)

func TestCrypto(t *testing.T) {
    data := []byte{0x01, 0x02, 0x03}
    key := [32]byte{
        0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
        0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
        0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
        0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08 }
    server, err := NewSecureSessionFromSharedKey(key)
    client, err := NewSecureClientSessionFromSharedKey(key)
    assert.Nil(t, err)
    
    var b bytes.Buffer
    b.Write(data)
    
    encrypted, err := server.Encrypt(&b)
    assert.Nil(t, err)
    
    decrypted, err := client.Decrypt(encrypted)
    assert.Nil(t, err)
    
    orig, err := ioutil.ReadAll(decrypted)
    assert.Nil(t, err)
    assert.Equal(t, orig, data)
}