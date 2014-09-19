package gohap

import(
    "crypto/sha512"
    "github.com/tonnerre/golang-go.crypto/hkdf"
    "io"
)

// Returns a 256-bit key
func HKDF_SHA512_256(master, salt, info []byte) ([]byte, error){
    hash := sha512.New
    hkdf := hkdf.New(hash, master, salt, info)
    
    key := make([]byte, 32) // 256 bit
    _, err := io.ReadFull(hkdf, key)
    
    return key, err
}