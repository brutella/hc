package crypto

import (
	_ "crypto/cipher"
	"encoding/binary"
	"encoding/hex"
	"github.com/codahale/chacha20"
	"github.com/tonnerre/golang-go.crypto/poly1305"

	"github.com/brutella/hap/common"
)

// Chacha20DecryptAndPoly1305Verify returns the chacha20 decrypted messages.
// An error is returned when the poly1305 message authenticator (seal) could not be verified.
// Nonce should be 8 byte.
func Chacha20DecryptAndPoly1305Verify(key, nonce, message []byte, mac [16]byte, add []byte) ([]byte, error) {

	chacha20, err := chacha20.NewCipher(key, nonce)
	if err != nil {
		panic(err)
	}

	// poly1305 key is chacha20 over 32 zeros
	var poly1305_key [32]byte
	var chacha20_key_out = make([]byte, 64)
	var zeros = make([]byte, 64)
	chacha20.XORKeyStream(chacha20_key_out, zeros)
	copy(poly1305_key[:], chacha20_key_out)

	var chacha20_out = make([]byte, len(message))
	var poly1305_out [16]byte

	// poly1305 byte order
	// - add bytes up to mod 16 (if available)
	// - message up to mod 16
	// - number of add bytes up to mod 8
	// - number of message bytes up to mod 8
	poly1305_in := make([]byte, 0)
	if len(add) > 0 {
		poly1305_in = AddBytes(poly1305_in, add, 16)
	}

	poly1305_in = AddBytes(poly1305_in, message, 16)
	add_len := make([]byte, 8)
	message_len := make([]byte, 8)
	binary.LittleEndian.PutUint64(add_len, uint64(len(add)))
	binary.LittleEndian.PutUint64(message_len, uint64(len(message)))

	poly1305_in = AddBytes(poly1305_in, add_len, 8)
	poly1305_in = AddBytes(poly1305_in, message_len, 8)

	poly1305.Sum(&poly1305_out, poly1305_in, &poly1305_key)

	if poly1305.Verify(&mac, poly1305_in, &poly1305_key) == false {
		return nil, common.NewError("MAC not equal: " + hex.EncodeToString(poly1305_out[:]) + " != " + hex.EncodeToString(mac[:]))
	}

	chacha20.XORKeyStream(chacha20_out, message)
	return chacha20_out, nil
}

// Chacha20EncryptAndPoly1305Seal returns the chacha20 encrypted message and poly1305 message authentictor (also refered as seals)
// Nonce should be 8 byte
func Chacha20EncryptAndPoly1305Seal(key, nonce, message []byte, add []byte) ([]byte /*encrypted*/, [16]byte /*mac*/, error) {

	chacha20, err := chacha20.NewCipher(key, nonce)
	if err != nil {
		panic(err)
	}

	// poly1305 key is chacha20 over 32 zeros
	var poly1305_key [32]byte
	var chacha20_key_out = make([]byte, 64)
	var zeros = make([]byte, 64)
	chacha20.XORKeyStream(chacha20_key_out, zeros)
	copy(poly1305_key[:], chacha20_key_out)

	var chacha20_out = make([]byte, len(message))
	var poly1305_out [16]byte
	chacha20.XORKeyStream(chacha20_out, message)

	poly1305_in := make([]byte, 0)
	if len(add) > 0 {
		poly1305_in = AddBytes(poly1305_in, add, 16)
	}

	poly1305_in = AddBytes(poly1305_in, chacha20_out, 16)
	add_len := make([]byte, 8)
	message_len := make([]byte, 8)
	binary.LittleEndian.PutUint64(add_len, uint64(len(add)))
	binary.LittleEndian.PutUint64(message_len, uint64(len(message)))

	poly1305_in = AddBytes(poly1305_in, add_len, 8)
	poly1305_in = AddBytes(poly1305_in, message_len, 8)

	poly1305.Sum(&poly1305_out, poly1305_in, &poly1305_key)

	return chacha20_out, poly1305_out, nil
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
