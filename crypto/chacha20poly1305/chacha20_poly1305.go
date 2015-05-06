package chacha20poly1305

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/codahale/chacha20"
	"golang.org/x/crypto/poly1305"
)

// DecryptAndVerify returns the chacha20 decrypted messages.
// An error is returned when the poly1305 message authenticator (seal) could not be verified.
// Nonce should be 8 byte.
func DecryptAndVerify(key, nonce, message []byte, mac [16]byte, add []byte) ([]byte, error) {

	chacha20, err := chacha20.New(key, nonce)
	if err != nil {
		panic(err)
	}

	// poly1305 key is chacha20 over 32 zeros
	var poly1305Key [32]byte
	var chacha20KeyOut = make([]byte, 64)
	var zeros = make([]byte, 64)
	chacha20.XORKeyStream(chacha20KeyOut, zeros)
	copy(poly1305Key[:], chacha20KeyOut)

	var chacha20Out = make([]byte, len(message))
	var poly1305Out [16]byte

	// poly1305 byte order
	// - add bytes up to mod 16 (if available)
	// - message up to mod 16
	// - number of add bytes up to mod 8
	// - number of message bytes up to mod 8
	var poly1305In []byte
	if len(add) > 0 {
		poly1305In = AddBytes(poly1305In, add, 16)
	}

	poly1305In = AddBytes(poly1305In, message, 16)
	addLength := make([]byte, 8)
	msgLength := make([]byte, 8)
	binary.LittleEndian.PutUint64(addLength, uint64(len(add)))
	binary.LittleEndian.PutUint64(msgLength, uint64(len(message)))

	poly1305In = AddBytes(poly1305In, addLength, 8)
	poly1305In = AddBytes(poly1305In, msgLength, 8)

	poly1305.Sum(&poly1305Out, poly1305In, &poly1305Key)

	if poly1305.Verify(&mac, poly1305In, &poly1305Key) == false {
		return nil, errors.New("MAC not equal: " + hex.EncodeToString(poly1305Out[:]) + " != " + hex.EncodeToString(mac[:]))
	}

	chacha20.XORKeyStream(chacha20Out, message)
	return chacha20Out, nil
}

// EncryptAndSeal returns the chacha20 encrypted message and poly1305 message authentictor (also refered as seals)
// Nonce should be 8 byte
func EncryptAndSeal(key, nonce, message []byte, add []byte) ([]byte /*encrypted*/, [16]byte /*mac*/, error) {

	chacha20, err := chacha20.New(key, nonce)
	if err != nil {
		panic(err)
	}

	// poly1305 key is chacha20 over 32 zeros
	var poly1305Key [32]byte
	var chacha20KeyOut = make([]byte, 64)
	var zeros = make([]byte, 64)
	chacha20.XORKeyStream(chacha20KeyOut, zeros)
	copy(poly1305Key[:], chacha20KeyOut)

	var chacha20Out = make([]byte, len(message))
	var poly1305Out [16]byte
	chacha20.XORKeyStream(chacha20Out, message)

	var poly1305In []byte
	if len(add) > 0 {
		poly1305In = AddBytes(poly1305In, add, 16)
	}

	poly1305In = AddBytes(poly1305In, chacha20Out, 16)
	addLength := make([]byte, 8)
	msgLength := make([]byte, 8)
	binary.LittleEndian.PutUint64(addLength, uint64(len(add)))
	binary.LittleEndian.PutUint64(msgLength, uint64(len(message)))

	poly1305In = AddBytes(poly1305In, addLength, 8)
	poly1305In = AddBytes(poly1305In, msgLength, 8)

	poly1305.Sum(&poly1305Out, poly1305In, &poly1305Key)

	return chacha20Out, poly1305Out, nil
}

// AddBytes appends *add* to *b*
// Additional bytes are appended to fill up until mod
//
// Example
//      b = []
//      add = [0xFF] -> [255]
//      mod = 8
//      result: [0xFF 0x0 0x0 0x0 0x0 0x0 0x0 0x0]
func AddBytes(b, add []byte, mod int) []byte {
	b = append(b, add...)
	if len(add)%mod != 0 {
		zeros := make([]byte, mod-len(add)%mod)
		b = append(b, zeros...)
	}

	return b
}
