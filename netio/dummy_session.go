package netio

import(
    "io"
)

type dummySession struct {
}

func NewPlainSession() Session {
    return &dummySession{}
}

// Session interface
func (s *dummySession) EncryptionEnabled() bool {
    return false
}

func (s *dummySession) Encrypt(r io.Reader) (io.Reader, error){
    return r, nil
}

func (s *dummySession) Decrypt(r io.Reader) (io.Reader, error){
    return r, nil
}