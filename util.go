package gohap

import(
    "encoding/binary"
    "bytes"
    "github.com/tang0th/go-chacha20"
    "github.com/tonnerre/golang-go.crypto/poly1305"
    "github.com/agl/ed25519"
    "fmt"
)

func DecryptAndVerify(key, nonce, message, mac, add []byte) ([]byte, error) {
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
    
    if bytes.Equal(poly1305_out[:], mac) == false {
        return nil, NewErrorf("MAC incorrect")
    }
    
    chacha20.XORKeyStream(chacha20_out, message, nonce, key)
    return chacha20_out, nil
}

// Encrypts and seals a message
// The returns values are the encrypted data, mac, error
func EncryptAndSeal(key, nonce, message []byte, mac [16]byte, add []byte) ([]byte /*encrypted*/, [16]byte /*mac*/, error) {
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

func ValidateED25519Signature(key, data, signature []byte) bool {
    if len(key) != ed25519.PublicKeySize || len(signature) != ed25519.SignatureSize {
        fmt.Printf("Invalid size of key (%d) or signature (%d)\n", len(key), len(signature))
        return false
    }
    
    var k [ed25519.PublicKeySize]byte
    var s [ed25519.SignatureSize]byte
    copy(k[:], key)
    copy(s[:], signature)
    
    return ed25519.Verify(&k, data, &s)
}

// Signs (ED25519) data based on public and secret key
// TODO can we just use public key as ed25519 sign key?
func ED25519Signature(key, data []byte) ([]byte, error) {
    if len(key) != ed25519.PrivateKeySize {
        return nil, NewErrorf("Invalid size of key (%d)\n", len(key))
    }
    
    var k [ed25519.PrivateKeySize]byte
    copy(k[:], key)
    signature := ed25519.Sign(&k, data)
    
    return signature[:], nil
}

func ED25519GenerateKey(str string) ([]byte/* public */, []byte /* secret */, error) {
    b := bytes.NewBuffer([]byte(str))
    if len(str) < 32 {
        zeros := make([]byte, 32 - len(str))
        b.Write(zeros)
    }
    
    public, secret, err := ed25519.GenerateKey(bytes.NewReader(b.Bytes()))
    
    return public[:], secret[:], err
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
