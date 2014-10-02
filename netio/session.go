package netio

import(
)

type Session interface {
    Decrypter() Decrypter
    Encrypter() Encrypter
    
    PairStartHandler() ContainerHandler
    PairVerifyHandler() PairVerifyHandler
    
    SetCryptographer(c Cryptographer)
    SetPairStartHandler(c ContainerHandler)
    SetPairVerifyHandler(c PairVerifyHandler)
}

type session struct {
    cryptographer Cryptographer
    pairStartHandler ContainerHandler
    pairVerifyHandler PairVerifyHandler
    
    nextCryptographer Cryptographer
}

func NewSession() *session {
    s := session{}
    
    return &s
}

func (s *session) Decrypter() Decrypter {
    if s.nextCryptographer != nil {
        s.cryptographer = s.nextCryptographer
        s.nextCryptographer = nil
    }
    
    return s.cryptographer
}

func (s *session) Encrypter() Encrypter {
    return s.cryptographer
}

func (s *session) PairStartHandler() ContainerHandler {
    return s.pairStartHandler
}

func (s *session) PairVerifyHandler() PairVerifyHandler {
    return s.pairVerifyHandler
}

func (s *session) SetCryptographer(c Cryptographer) {
    s.nextCryptographer = c
}
func (s *session) SetPairStartHandler(c ContainerHandler) {
    s.pairStartHandler = c
}

func (s *session) SetPairVerifyHandler(c PairVerifyHandler) {
    s.pairVerifyHandler = c    
}