package crypto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/brutella/hc/crypto/chacha20poly1305"
	"github.com/brutella/hc/crypto/hkdf"
	"io"
	"sync/atomic"
)

// secureSession provide a secure session by encrypting and decrypting data
type secureSession struct {
	encryptKey [32]byte
	decryptKey [32]byte

	encryptCount uint64
	decryptCount uint64

	readEncrypted bool
}

// NewSecureSessionFromSharedKey returns a session from a shared private key.
func NewSecureSessionFromSharedKey(sharedKey [32]byte) (Cryptographer, error) {
	salt := []byte("Control-Salt")
	out := []byte("Control-Read-Encryption-Key")
	in := []byte("Control-Write-Encryption-Key")

	var s = new(secureSession)
	var err error
	s.encryptKey, err = hkdf.Sha512(sharedKey[:], salt, out)
	s.encryptCount = 0
	if err != nil {
		return nil, err
	}

	s.decryptKey, err = hkdf.Sha512(sharedKey[:], salt, in)
	s.decryptCount = 0

	return s, err
}

// NewSecureClientSessionFromSharedKey returns a session from a shared secret key to simulate a HomeKit client.
// This is currently only used for testing.
func NewSecureClientSessionFromSharedKey(sharedKey [32]byte) (Cryptographer, error) {
	salt := []byte("Control-Salt")
	out := []byte("Control-Write-Encryption-Key")
	in := []byte("Control-Read-Encryption-Key")

	var s = new(secureSession)
	var err error
	s.encryptKey, err = hkdf.Sha512(sharedKey[:], salt, out)
	s.encryptCount = 0
	if err != nil {
		return nil, err
	}

	s.decryptKey, err = hkdf.Sha512(sharedKey[:], salt, in)
	s.decryptCount = 0

	return s, err
}

// Encrypt return the encrypted data by splitting it into packets
// [ length (2 bytes)] [ data ] [ auth (16 bytes)]
func (s *secureSession) Encrypt(r io.Reader) (io.Reader, error) {
	packets := packetsFromBytes(r)
	var buf bytes.Buffer
	for _, p := range packets {
		var nonce [8]byte
		binary.LittleEndian.PutUint64(nonce[:], atomic.AddUint64(&s.encryptCount, 1)-1)

		bLength := make([]byte, 2)
		binary.LittleEndian.PutUint16(bLength, uint16(p.length))

		encrypted, mac, err := chacha20poly1305.EncryptAndSeal(s.encryptKey[:], nonce[:], p.value, bLength[:])
		if err != nil {
			return nil, err
		}

		buf.Write(bLength[:])
		buf.Write(encrypted)
		buf.Write(mac[:])
	}

	return &buf, nil
}

// Decrypt returns the decrypted data
func (s *secureSession) Decrypt(r io.Reader) (io.Reader, error) {
	var buf bytes.Buffer
	for {
		var length uint16
		if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if length > PacketLengthMax {
			return nil, fmt.Errorf("Packet size too big %d", length)
		}

		var b = make([]byte, length)
		if err := binary.Read(r, binary.LittleEndian, &b); err != nil {
			return nil, err
		}

		var mac [16]byte
		if err := binary.Read(r, binary.LittleEndian, &mac); err != nil {
			return nil, err
		}

		var nonce [8]byte
		binary.LittleEndian.PutUint64(nonce[:], atomic.AddUint64(&s.decryptCount, 1)-1)

		lengthBytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(lengthBytes, uint16(length))

		decrypted, err := chacha20poly1305.DecryptAndVerify(s.decryptKey[:], nonce[:], b, mac, lengthBytes)

		if err != nil {
			return nil, fmt.Errorf("Data encryption failed %s", err)
		}

		buf.Write(decrypted)

		// Finish when all bytes fit in b
		if length < PacketLengthMax {
			break
		}
	}

	return &buf, nil
}
