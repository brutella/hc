package netio

import(
)

type Session interface {
    // For decrypting incoming data, may be nil
    Decrypter() Decrypter
    
    // For encrypting outgoing data, may be nil
    Encrypter() Encrypter
    
    PairSetupHandler() ContainerHandler
    PairVerifyHandler() PairVerifyHandler
    
    SetCryptographer(c Cryptographer)
    SetPairSetupHandler(c ContainerHandler)
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
    // Return the next cryptographer when possible
    // This allows sessions to switch encryption
    if s.nextCryptographer != nil {
        s.cryptographer = s.nextCryptographer
        s.nextCryptographer = nil
    }
    
    return s.cryptographer
}

func (s *session) Encrypter() Encrypter {
    return s.cryptographer
}

func (s *session) PairSetupHandler() ContainerHandler {
    return s.pairStartHandler
}

func (s *session) PairVerifyHandler() PairVerifyHandler {
    return s.pairVerifyHandler
}

func (s *session) SetCryptographer(c Cryptographer) {
    s.nextCryptographer = c
}
func (s *session) SetPairSetupHandler(c ContainerHandler) {
    s.pairStartHandler = c
}

func (s *session) SetPairVerifyHandler(c PairVerifyHandler) {
    s.pairVerifyHandler = c    
}