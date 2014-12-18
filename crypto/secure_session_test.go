package crypto

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "bytes"
    "io/ioutil"
    "math"
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
    
    // Set count to min 2 bytes to test byte order handling
    server.encryptCount = 128
    encrypted, err := server.Encrypt(&b)
    assert.Nil(t, err)

    client.decryptCount = 128
    decrypted, err := client.Decrypt(encrypted)
    assert.Nil(t, err)
    orig, err := ioutil.ReadAll(decrypted)
    assert.Nil(t, err)
    assert.Equal(t, orig, data)
}

func TestCryptoMaxPacketCount(t *testing.T) {
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
    server.encryptCount = math.MaxUint64
    encrypted, err := server.Encrypt(&b)
    assert.Nil(t, err)
    assert.Equal(t, server.encryptCount, uint64(0))

    client.decryptCount = math.MaxUint64
    decrypted, err := client.Decrypt(encrypted)
    assert.Nil(t, err)
    assert.Equal(t, client.decryptCount, uint64(0))
    orig, err := ioutil.ReadAll(decrypted)
    assert.Nil(t, err)
    assert.Equal(t, orig, data)
}