package gohap

import(
    "crypto/sha512"
    // "github.com/tadglines/go-pkgs/crypto/srp"
    "github.com/theojulienne/go-srp/crypto/srp"
)

func SRP6Password(username, password []byte) []byte {
    // H(U | ":" | P)
    h := sha512.New()
    h.Write(username)
    h.Write([]byte(":"))
    return h.Sum(password)
}
// x = H(s | H(I | ":" | P))
func SHA512KeyDerivativeFunction(username []byte) srp.KeyDerivationFunc {
	return func(salt, password []byte) []byte {
        h := sha512.New()
        h.Write(username)
        h.Write([]byte(":"))
        h.Write(password)
        t2 := h.Sum(nil)
        h.Reset()
        h.Write(salt)
        h.Write(t2)
        return h.Sum(nil)
	}
}
