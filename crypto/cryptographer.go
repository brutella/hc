package crypto

import (
	"io"
)

// Encrypter encrypts bytes.
type Encrypter interface {
	Encrypt(r io.Reader) (io.Reader, error)
}

// Decrypter decrypts bytes.
type Decrypter interface {
	Decrypt(r io.Reader) (io.Reader, error)
}

// A Cryptographer is a De- and Encrypter.
type Cryptographer interface {
	Encrypter
	Decrypter
}
