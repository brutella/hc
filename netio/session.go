package netio

import (
	"net"
)

type Session interface {
	// For decrypting incoming data, may be nil
	Decrypter() Decrypter

	// For encrypting outgoing data, may be nil
	Encrypter() Encrypter
	// Sets the cryptographer for encryption and decryption
	SetCryptographer(c Cryptographer)

	PairSetupHandler() ContainerHandler
	PairVerifyHandler() PairVerifyHandler
	SetPairSetupHandler(c ContainerHandler)
	SetPairVerifyHandler(c PairVerifyHandler)

	Connection() net.Conn
}

type session struct {
	cryptographer     Cryptographer
	pairStartHandler  ContainerHandler
	pairVerifyHandler PairVerifyHandler
	connection        net.Conn

	nextCryptographer Cryptographer
}

func NewSession(connection net.Conn) *session {
	s := session{
		connection: connection,
	}

	return &s
}

func (s *session) Connection() net.Conn {
	return s.connection
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
