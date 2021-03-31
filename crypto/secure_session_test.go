package crypto

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
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
	if err != nil {
		t.Fatal(err)
	}

	client, err := NewSecureClientSessionFromSharedKey(key)
	if err != nil {
		t.Fatal(err)
	}

	// Set count to min 2 bytes to test byte order handling
	secServer := server.(*secureSession)
	secServer.encryptCount = 128

	secClient := client.(*secureSession)
	secClient.decryptCount = 128

	for i := 0; i < 2; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			var b bytes.Buffer
			b.Write(data)

			encrypted, err := server.Encrypt(&b)
			if err != nil {
				t.Fatal(err)
			}

			decrypted, err := client.Decrypt(encrypted)
			if err != nil {
				t.Fatal(err)
			}

			orig, err := ioutil.ReadAll(decrypted)
			if err != nil {
				t.Fatal(err)
			}
			if reflect.DeepEqual(orig, data) == false {
				t.Fatal("invalid decryption")
			}
		})
	}
}

func TestCryptoMaxPacketCount(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	key := [32]byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	server, err := NewSecureSessionFromSharedKey(key)
	if err != nil {
		t.Fatal(err)
	}

	client, err := NewSecureClientSessionFromSharedKey(key)
	if err != nil {
		t.Fatal(err)
	}

	var b bytes.Buffer
	b.Write(data)
	secServer := server.(*secureSession)
	secServer.encryptCount = math.MaxUint64
	encrypted, err := server.Encrypt(&b)
	if err != nil {
		t.Fatal(err)
	}

	if secServer.encryptCount != 0 {
		t.Fatal(secServer.encryptCount)
	}

	secClient := client.(*secureSession)
	secClient.decryptCount = math.MaxUint64
	decrypted, err := client.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	if secClient.decryptCount != 0 {
		t.Fatal(secServer.encryptCount)
	}
	orig, err := ioutil.ReadAll(decrypted)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(orig, data) == false {
		t.Fatal("invalid decryption")
	}
}

func TestCryptoMaxPacketLength(t *testing.T) {
	// 1044 bytes
	data := []byte("1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz")
	if len(data) != 1044 {
		t.Fatal(len(data))
	}

	key := [32]byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	server, err := NewSecureSessionFromSharedKey(key)
	if err != nil {
		t.Fatal(err)
	}

	client, err := NewSecureClientSessionFromSharedKey(key)
	if err != nil {
		t.Fatal(err)
	}

	var b bytes.Buffer
	b.Write(data)
	encrypted, err := server.Encrypt(&b)
	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := client.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	orig, err := ioutil.ReadAll(decrypted)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(orig, data) == false {
		t.Fatal("invalid decryption")
	}
}
