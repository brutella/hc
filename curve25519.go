package gohap

import (
    "crypto/rand"
    "github.com/tonnerre/golang-go.crypto/curve25519"
)

const (
    KeySize = 32
)

func Curve25519_GenerateSecretKey() [KeySize]byte {
    var b [KeySize]byte
    rand.Read(b[:])
    
    return b
}

func Curve25519_PublicKey(secretKey [KeySize]byte) [KeySize]byte {
    var k [KeySize]byte
    curve25519.ScalarBaseMult(&k, &secretKey)
    
    return k
}

func Curve25519_SharedSecret(secretKey, otherPublicKey [KeySize]byte) [KeySize]byte {
    var k [KeySize]byte
    curve25519.ScalarMult(&k, &secretKey, &otherPublicKey)
    
    return k
}