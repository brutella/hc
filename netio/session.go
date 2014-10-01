package netio

import(
    "io"
)

type Session interface {
    Encrypt(r io.Reader) (io.Reader, error);
    Decrypt(r io.Reader) (io.Reader, error);
    EncryptionEnabled() bool
}