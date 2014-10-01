package crypto

import(
    "io"
    "crypto/sha512"
    "github.com/tonnerre/golang-go.crypto/hkdf"
)

// Returns a 256-bit key
func HKDF_SHA512(master, salt, info []byte) ([32]byte, error){
    hash := sha512.New
    hkdf := hkdf.New(hash, master, salt, info)
    
    key := make([]byte, 32) // 256 bit
    _, err := io.ReadFull(hkdf, key)
    
    var result [32]byte
    copy(result[:], key)
    
    return result, err
}