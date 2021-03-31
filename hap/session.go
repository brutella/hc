package hap

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/crypto"
	"net"
	"sync"
)

// Session contains objects (encrypter, decrypter, pairing handler,...) used to handle the data communication.
type Session interface {
	// Decrypter returns decrypter for incoming data, may be nil
	Decrypter() crypto.Decrypter

	// Encrypter returns encrypter for outgoing data, may be nil
	Encrypter() crypto.Encrypter

	// SetCryptographer sets the new cryptographer used for en-/decryption
	SetCryptographer(c crypto.Cryptographer)

	// PairSetupHandler returns the pairing setup handler
	PairSetupHandler() ContainerHandler

	// SetPairSetupHandler sets the handler for pairing setup
	SetPairSetupHandler(c ContainerHandler)

	// PairVerifyHandler returns the pairing verify handler
	PairVerifyHandler() PairVerifyHandler

	// SetPairVerifyHandler sets the handler for pairing verify
	SetPairVerifyHandler(c PairVerifyHandler)

	// Connection returns the associated connection
	Connection() net.Conn

	Subscribe(*characteristic.Characteristic)
	Unsubscribe(*characteristic.Characteristic)
	IsSubscribedTo(*characteristic.Characteristic) bool
}

type session struct {
	cryptographer     crypto.Cryptographer
	pairStartHandler  ContainerHandler
	pairVerifyHandler PairVerifyHandler
	connection        net.Conn
	mu                *sync.RWMutex
	subs              map[*characteristic.Characteristic]bool

	// Temporary variable to reference next cryptographer
	nextCryptographer crypto.Cryptographer
}

// NewSession returns a session for a connection.
func NewSession(connection net.Conn) Session {
	s := session{
		connection: connection,
		mu:         &sync.RWMutex{},
		subs:       map[*characteristic.Characteristic]bool{},
	}

	return &s
}

func (s *session) Connection() net.Conn {
	return s.connection
}

func (s *session) Decrypter() crypto.Decrypter {
	// Return the next cryptographer when possible
	// This allows sessions to switch encryption
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.nextCryptographer != nil {
		s.cryptographer = s.nextCryptographer
		s.nextCryptographer = nil
	}

	return s.cryptographer
}

func (s *session) Encrypter() crypto.Encrypter {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.cryptographer
}

func (s *session) PairSetupHandler() ContainerHandler {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.pairStartHandler
}

func (s *session) PairVerifyHandler() PairVerifyHandler {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.pairVerifyHandler
}

func (s *session) SetCryptographer(c crypto.Cryptographer) {
	// Temporarily set the cryptographer as the nextCryptographer
	// The nextCryptographer is used the next time Decrypter() is called.
	// Otherwise the Encrypter() encrypts differently than the previous Decrypter()
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextCryptographer = c
}
func (s *session) SetPairSetupHandler(c ContainerHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pairStartHandler = c
}

func (s *session) SetPairVerifyHandler(c PairVerifyHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pairVerifyHandler = c
}

func (s *session) IsSubscribedTo(ch *characteristic.Characteristic) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.subs[ch] == true
}

func (s *session) Subscribe(ch *characteristic.Characteristic) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subs[ch] = true
}

func (s *session) Unsubscribe(ch *characteristic.Characteristic) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subs[ch] = false
}
