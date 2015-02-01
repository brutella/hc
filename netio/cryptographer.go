package netio

import (
	"io"
)

type Encrypter interface {
	Encrypt(r io.Reader) (io.Reader, error)
}

type Decrypter interface {
	Decrypt(r io.Reader) (io.Reader, error)
}

type Cryptographer interface {
	Encrypter
	Decrypter
}
