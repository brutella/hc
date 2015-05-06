package netio

import (
	"bytes"
	"github.com/brutella/hc/crypto"
	"github.com/brutella/log"
	"net"
	"time"

	"bufio"
	"io"
	"io/ioutil"
)

// HAPConnection is a connection connection based on HAP protocol which encrypts and decrypts the data.
//
// For every connection, a new session is created in the context. The session uses the Cryptographer
// to encrypt and decrypt data. The Cryptographer is created by one of the endpoint handlers after
// pairing has been verified. After that the communication is encrypted.
//
// When the connection is closed, the related session is removed from the context.
type HAPConnection struct {
	connection net.Conn
	context    HAPContext

	// Used to buffer reads
	readBuffer io.Reader
}

// NewHAPConnection returns a hap connection.
func NewHAPConnection(connection net.Conn, context HAPContext) *HAPConnection {
	conn := &HAPConnection{
		connection: connection,
		context:    context,
	}

	// Setup new session for the connection
	session := NewSession(conn)
	context.SetSessionForConnection(session, conn)

	return conn
}

// EncryptedWrite encrypts and writes bytes to the connection.
// The method returns the number of written bytes and an error when writing failed.
func (con *HAPConnection) EncryptedWrite(b []byte) (int, error) {
	var buffer bytes.Buffer
	buffer.Write(b)
	encrypted, err := con.getEncrypter().Encrypt(&buffer)

	if err != nil {
		log.Println("[ERRO] Encryption failed:", err)
		err = con.connection.Close()
		return 0, err
	}

	encryptedBytes, err := ioutil.ReadAll(encrypted)
	n, err := con.connection.Write(encryptedBytes)

	return n, err
}

// DecryptedRead reads and decrypts bytes from the connection.
// The method returns the number of read bytes and an error when reading failed.
func (con *HAPConnection) DecryptedRead(b []byte) (int, error) {
	if con.readBuffer == nil {
		buffered := bufio.NewReader(con.connection)
		decrypted, err := con.getDecrypter().Decrypt(buffered)
		if err != nil {
			log.Println("[ERRO] Decryption failed:", err)
			err = con.connection.Close()
			return 0, err
		}

		con.readBuffer = decrypted
	}

	n, err := con.readBuffer.Read(b)

	if n < len(b) || err == io.EOF {
		con.readBuffer = nil
	}

	return n, err
}

// Write writes bytes to the connection.
// The written bytes are encrypted when possible.
func (con *HAPConnection) Write(b []byte) (int, error) {
	if con.getEncrypter() != nil {
		return con.EncryptedWrite(b)
	}

	return con.connection.Write(b)
}

// Read reads bytes from the connection. The read bytes are decrypted when possible.
func (con *HAPConnection) Read(b []byte) (int, error) {
	if con.getDecrypter() != nil {
		return con.DecryptedRead(b)
	}

	return con.connection.Read(b)
}

// Close closes the connection and deletes the related session from the context.
func (con *HAPConnection) Close() error {
	log.Println("[INFO] Close connection and remove session")

	// Remove session from the context
	con.context.DeleteSessionForConnection(con.connection)

	return con.connection.Close()
}

// LocalAddr calls LocalAddr() of the underlying connection
func (con *HAPConnection) LocalAddr() net.Addr {
	return con.connection.LocalAddr()
}

// RemoteAddr calls RemoteAddr() of the underlying connection
func (con *HAPConnection) RemoteAddr() net.Addr {
	return con.connection.RemoteAddr()
}

// SetDeadline calls SetDeadline() of the underlying connection
func (con *HAPConnection) SetDeadline(t time.Time) error {
	return con.connection.SetReadDeadline(t)
}

// SetReadDeadline calls SetReadDeadline() of the underlying connection
func (con *HAPConnection) SetReadDeadline(t time.Time) error {
	return con.connection.SetReadDeadline(t)
}

// SetWriteDeadline calls SetWriteDeadline() of the underlying connection
func (con *HAPConnection) SetWriteDeadline(t time.Time) error {
	return con.connection.SetWriteDeadline(t)
}

// getEncrypter returns the session's Encrypter, otherwise nil
func (con *HAPConnection) getEncrypter() crypto.Encrypter {
	session := con.context.GetSessionForConnection(con.connection)
	if session != nil {
		return session.Encrypter()
	}

	return nil
}

// getDecrypter returns the session's Decrypter, otherwise nil
func (con *HAPConnection) getDecrypter() crypto.Decrypter {
	session := con.context.GetSessionForConnection(con.connection)
	if session != nil {
		return session.Decrypter()
	}

	return nil
}
