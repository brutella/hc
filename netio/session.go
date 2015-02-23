package netio

import (
	"net"
)

type Session interface {
	// Decrypter returns decrypter for incoming data, may be nil
	Decrypter() Decrypter

	// Encrypter returns encrypter for outgoing data, may be nil
	Encrypter() Encrypter

	// SetCryptographer sets the new cryptographer used for en-/decryption
	SetCryptographer(c Cryptographer)

	// PairSetupHandler returns the pairing setup handler
	PairSetupHandler() ContainerHandler

	// SetPairSetupHandler sets the handler for pairing setup
	SetPairSetupHandler(c ContainerHandler)

	// PairVerifyHandler returns the pairing verify handler
	PairVerifyHandler() PairVerifyHandler

	// SetPairVerifyHandler sets the handler for pairing verify
	SetPairVerifyHandler(c PairVerifyHandler)

	// Returns the associated connection
	Connection() net.Conn
}

type session struct {
	cryptographer     Cryptographer
	pairStartHandler  ContainerHandler
	pairVerifyHandler PairVerifyHandler
	connection        net.Conn

	// Temporary variable to reference next cryptographer
	nextCryptographer Cryptographer
}

// NewSession creates a new session for a connection
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
	// Temporarily set the cryptographer as the nextCryptographer
	// The nextCryptographer is used the next time Decrypter() is called.
	// Otherwise the Encrypter() encrypts differently than the previous Decrypter()
	s.nextCryptographer = c
}
func (s *session) SetPairSetupHandler(c ContainerHandler) {
	s.pairStartHandler = c
}

func (s *session) SetPairVerifyHandler(c PairVerifyHandler) {
	s.pairVerifyHandler = c
}
