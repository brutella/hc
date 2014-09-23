package gohap

import(
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
)

// x = SHA1(s | SHA1(I | ":" | P))
func SHA512KeyDerivativeFunction(username []byte) srp.KeyDerivationFunc {
	return func(salt, password []byte) []byte {
        h := sha512.New()
        h.Write(username)
        h.Write([]byte(":"))
        result := h.Sum(password)
        
        h.Reset()
        h.Write(salt)
        return h.Sum(result)
	}
}
