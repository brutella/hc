package gohap

import(
    "encoding/binary"
    "encoding/hex"
    "bytes"
    "github.com/tang0th/go-chacha20"
    "github.com/tonnerre/golang-go.crypto/poly1305"
)

// Decrypts a message with chacha20 and verifies it with poly1305
func Chacha20DecryptAndPoly1305Verify(key, nonce, message []byte, mac [16]byte, add []byte) ([]byte, error) {
    var chacha20_out = make([]byte, len(message))
    var poly1305_out [16]byte
    var poly1305_key [32]byte
    var zeros = make([]byte, 32)
    
    // poly1305 key is chacha20 over 32 zeros
    chacha20.XORKeyStream(chacha20_out, zeros, nonce, key)
    copy(chacha20_out[:], poly1305_key[:])
    
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
    
    poly1305_in = AddBytes(poly1305_in, add_len ,8)
    poly1305_in = AddBytes(poly1305_in, message_len, 8)
    
    poly1305.Sum(&poly1305_out, poly1305_in, &poly1305_key)
    
    if bytes.Equal(poly1305_out[:], mac[:]) == false {
        return nil, NewErrorf("MAC incorrect %s", hex.EncodeToString(poly1305_out[:]), hex.EncodeToString(mac[:]))
    }
    
    chacha20.XORKeyStream(chacha20_out, message, nonce, key)
    return chacha20_out, nil
}

// Encrypts the message with chacha20 and seals the message with poly1305
func Chacha20EncryptAndPoly1305Seal(key, nonce, message []byte, mac [16]byte, add []byte) ([]byte /*encrypted*/, [16]byte /*mac*/, error) {
    var chacha20_out = make([]byte, len(message))
    var poly1305_out [16]byte
    var poly1305_key [32]byte
    var zeros = make([]byte, 32)
    
    // poly1305 key is chacha20 over 32 zeros
    chacha20.XORKeyStream(chacha20_out, zeros, nonce, key)
    copy(chacha20_out[:], poly1305_key[:])
    
    chacha20.XORKeyStream(chacha20_out, message, nonce, key)
    
    poly1305_in := make([]byte, 0)
    if len(add) > 0 {
        poly1305_in = AddBytes(poly1305_in, add, 16)
    }
    
    poly1305_in = AddBytes(poly1305_in, message, 16)
    add_len := make([]byte, 8)
    message_len := make([]byte, 8)
    binary.LittleEndian.PutUint64(add_len, uint64(len(add)))
    binary.LittleEndian.PutUint64(message_len, uint64(len(message)))
    
    poly1305_in = AddBytes(poly1305_in, add_len ,8)
    poly1305_in = AddBytes(poly1305_in, message_len, 8)
    
    poly1305.Sum(&poly1305_out, poly1305_in, &poly1305_key)
    
    return chacha20_out, poly1305_out, nil
}

// Appends `add` to `b`
// Additional bytes are appended to fill up until mod
// 
// Example
// b = []
// add = [0xFF] -> [255]
// mod = 8
// result: [0xFF 0x0 0x0 0x0 0x0 0x0 0x0 0x0]
func AddBytes(b, add []byte, mod int) []byte {
    b = append(b, add...)
    if len(add) % mod != 0 {
        zeros := make([]byte, mod - len(add) % mod)
        b = append(b, zeros...)
    }
    
    return b
}
