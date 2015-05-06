package crypto

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math"
	"testing"
)

func TestCrypto(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	key := [32]byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	server, err := NewSecureSessionFromSharedKey(key)
	client, err := NewSecureClientSessionFromSharedKey(key)
	assert.Nil(t, err)

	var b bytes.Buffer
	b.Write(data)

	// Set count to min 2 bytes to test byte order handling
	secServer := server.(*secureSession)
	secServer.encryptCount = 128
	encrypted, err := server.Encrypt(&b)
	assert.Nil(t, err)

	secClient := client.(*secureSession)
	secClient.decryptCount = 128
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
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	server, err := NewSecureSessionFromSharedKey(key)
	client, err := NewSecureClientSessionFromSharedKey(key)
	assert.Nil(t, err)

	var b bytes.Buffer
	b.Write(data)
	secServer := server.(*secureSession)
	secServer.encryptCount = math.MaxUint64
	encrypted, err := server.Encrypt(&b)
	assert.Nil(t, err)
	assert.Equal(t, secServer.encryptCount, uint64(0))

	secClient := client.(*secureSession)
	secClient.decryptCount = math.MaxUint64
	decrypted, err := client.Decrypt(encrypted)
	assert.Nil(t, err)
	assert.Equal(t, secClient.decryptCount, uint64(0))
	orig, err := ioutil.ReadAll(decrypted)
	assert.Nil(t, err)
	assert.Equal(t, orig, data)
}

func TestCryptoMaxPacketLength(t *testing.T) {
	// 1044 bytes
	data := []byte("1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz")
	assert.Equal(t, len(data), 1044)

	key := [32]byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
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
