package hap

import (
	"bytes"
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/log"
	"net"
	"time"

	"bufio"
	"io"
	"io/ioutil"
)

// Connection is a connection based on HAP protocol which encrypts and decrypts the data.
//
// For every connection, a new session is created in the context. The session uses the Cryptographer
// to encrypt and decrypt data. The Cryptographer is created by one of the endpoint handlers after
// pairing has been verified. After that the communication is encrypted.
//
// When the connection is closed, the related session is removed from the context.
type Connection struct {
	connection net.Conn
	context    Context

	// Used to buffer reads
	readBuffer io.Reader
}

// NewConnection returns a hap connection.
func NewConnection(connection net.Conn, context Context) *Connection {
	conn := &Connection{
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
func (con *Connection) EncryptedWrite(b []byte) (int, error) {
	var buffer bytes.Buffer
	buffer.Write(b)
	encrypted, err := con.getEncrypter().Encrypt(&buffer)

	if err != nil {
		log.Info.Panic("Encryption failed:", err)
		err = con.connection.Close()
		return 0, err
	}

	encryptedBytes, err := ioutil.ReadAll(encrypted)
	n, err := con.connection.Write(encryptedBytes)

	return n, err
}

// DecryptedRead reads and decrypts bytes from the connection.
// The method returns the number of read bytes and an error when reading failed.
func (con *Connection) DecryptedRead(b []byte) (int, error) {
	if con.readBuffer == nil {
		buffered := bufio.NewReader(con.connection)
		decrypted, err := con.getDecrypter().Decrypt(buffered)
		if err != nil {
			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				// Ignore timeout error #77
			} else {
				log.Debug.Println("Decryption failed:", err)
				err = con.connection.Close()
			}
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
func (con *Connection) Write(b []byte) (int, error) {
	if con.getEncrypter() != nil {
		return con.EncryptedWrite(b)
	}

	return con.connection.Write(b)
}

// Read reads bytes from the connection. The read bytes are decrypted when possible.
func (con *Connection) Read(b []byte) (int, error) {
	if con.getDecrypter() != nil {
		return con.DecryptedRead(b)
	}

	return con.connection.Read(b)
}

// Close closes the connection and deletes the related session from the context.
func (con *Connection) Close() error {
	log.Debug.Println("Close connection and remove session")

	// Remove session from the context
	con.context.DeleteSessionForConnection(con.connection)

	return con.connection.Close()
}

// LocalAddr calls LocalAddr() of the underlying connection
func (con *Connection) LocalAddr() net.Addr {
	return con.connection.LocalAddr()
}

// RemoteAddr calls RemoteAddr() of the underlying connection
func (con *Connection) RemoteAddr() net.Addr {
	return con.connection.RemoteAddr()
}

// SetDeadline calls SetDeadline() of the underlying connection
func (con *Connection) SetDeadline(t time.Time) error {
	return con.connection.SetDeadline(t)
}

// SetReadDeadline calls SetReadDeadline() of the underlying connection
func (con *Connection) SetReadDeadline(t time.Time) error {
	return con.connection.SetReadDeadline(t)
}

// SetWriteDeadline calls SetWriteDeadline() of the underlying connection
func (con *Connection) SetWriteDeadline(t time.Time) error {
	return con.connection.SetWriteDeadline(t)
}

// getEncrypter returns the session's Encrypter, otherwise nil
func (con *Connection) getEncrypter() crypto.Encrypter {
	session := con.context.GetSessionForConnection(con.connection)
	if session != nil {
		return session.Encrypter()
	}

	return nil
}

// getDecrypter returns the session's Decrypter, otherwise nil
func (con *Connection) getDecrypter() crypto.Decrypter {
	session := con.context.GetSessionForConnection(con.connection)
	if session != nil {
		return session.Decrypter()
	}

	return nil
}
