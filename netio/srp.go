package netio

import(
    "github.com/tadglines/go-pkgs/crypto/srp"
)

// Main SRP algorithm is described in http://srp.stanford.edu/design.html
// The HAP uses the SRP-6a Stanford implementation with the following characteristics
//      x = H(s | H(I | ":" | P)) -> called the key derivate function
//      M1 = H(H(N) xor H(g), H(I), s, A, B, K)
const (
    SRPGroup = "rfc5054.3072" // N (modulo) => 384 byte
)

// Does x = H(s | H(I | ":" | P)) which is required by the Stanford SRP-6a implementation
func KeyDerivativeFuncRFC2945(h srp.HashFunc, username []byte) srp.KeyDerivationFunc {
	return func(salt, password []byte) []byte {
        h := h()
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
