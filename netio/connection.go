package netio

import (
	"bytes"
	"github.com/brutella/log"
	"net"
	"time"

	"bufio"
	"io"
	"io/ioutil"
)

// TCP connection based on HAP protocol
//
// The connections creates a new session in the context to support simultaneous encrypted connections
// with different encryption keys.
//
// The sessions' Encrypter and Decrypter are used to encrypt and decrypt data if possible. If no
// Encrypter and Decrypter are availabe, the connection is handles as plain.
//
// When the connection is closed, the session is removed from the context.
type hapConnection struct {
	connection net.Conn
	context    HAPContext

	// Used to buffer reads
	readBuffer io.Reader
}

// NewHAPConnection returns a new hap connection
// A new session is created and stored in the context for the new connection.
// The session holds references to Encrypter and Decrypter.
func NewHAPConnection(connection net.Conn, context HAPContext) *hapConnection {
	conn := &hapConnection{
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
func (con *hapConnection) EncryptedWrite(b []byte) (int, error) {
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
func (con *hapConnection) DecryptedRead(b []byte) (int, error) {
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
func (con *hapConnection) Write(b []byte) (int, error) {
	if con.getEncrypter() != nil {
		return con.EncryptedWrite(b)
	}

	return con.connection.Write(b)
}

// Read reads bytes from the connection.
// The read bytes are decrypted when possible.
func (con *hapConnection) Read(b []byte) (int, error) {
	if con.getDecrypter() != nil {
		return con.DecryptedRead(b)
	}

	return con.connection.Read(b)
}

// Close closes the connection and deletes the session from
// the context.
func (con *hapConnection) Close() error {
	log.Println("[INFO] Close connection and remove session")

	// Remove session from the context
	con.context.DeleteSessionForConnection(con.connection)

	return con.connection.Close()
}

func (con *hapConnection) LocalAddr() net.Addr {
	return con.connection.LocalAddr()
}

func (con *hapConnection) RemoteAddr() net.Addr {
	return con.connection.RemoteAddr()
}

func (con *hapConnection) SetDeadline(t time.Time) error {
	return con.connection.SetReadDeadline(t)
}

func (con *hapConnection) SetReadDeadline(t time.Time) error {
	return con.connection.SetReadDeadline(t)
}

func (con *hapConnection) SetWriteDeadline(t time.Time) error {
	return con.connection.SetWriteDeadline(t)
}

// getEncrypter returns the session's Encrypter, otherwise nil
func (c *hapConnection) getEncrypter() Encrypter {
	session := c.context.GetSessionForConnection(c.connection)
	if session != nil {
		return session.Encrypter()
	}

	return nil
}

// getDecrypter returns the session's Decrypter, otherwise nil
func (c *hapConnection) getDecrypter() Decrypter {
	session := c.context.GetSessionForConnection(c.connection)
	if session != nil {
		return session.Decrypter()
	}

	return nil
}
