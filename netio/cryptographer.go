package netio

import(
    "io"
    "fmt"
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

type dummyCryptographer struct {
}

func NewDummyCryptographer() *dummyCryptographer {
    return &dummyCryptographer{}
}

func (s *dummyCryptographer) Encrypt(r io.Reader) (io.Reader, error) {
    fmt.Println("Dummy Encrypt")
    return r, nil
}

func (s *dummyCryptographer) Decrypt(r io.Reader) (io.Reader, error) {
    fmt.Println("Dummy Decrypt")
    return r, nil
}

type cryptographer struct {
}

func NewCryptographer() *cryptographer {
    return &cryptographer{}
}

func (s *cryptographer) Encrypt(r io.Reader) (io.Reader, error) {
    fmt.Println("Secure Encrypt")
    return r, nil
}

func (s *cryptographer) Decrypt(r io.Reader) (io.Reader, error) {
    fmt.Println("Secure Decrypt")
    return r, nil
}