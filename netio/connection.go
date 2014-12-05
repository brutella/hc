package netio

import(
    "github.com/brutella/log"
    "time"
    "net"
    "bytes"
    
    "io/ioutil"
    "io"
    "bufio"
)

// TCP connection based on HAP protocol
//
// The connections creates a new session in the context to support simultaneous encrypted connections.
//
// The sessions' Encrypter and Decrypter are used to encrypt and decrypt data if possible. If no
// Encrypter and Decrypter are availabe, the connection is handles as plain.
//
// When the connection is closed, the session is removed from the context.
type tcpHAPConnection struct {
    connection net.Conn
    context HAPContext
    
    // Used to buffer reads
    readBuffer io.Reader
}

func NewHAPConnection(connection net.Conn, context HAPContext) *tcpHAPConnection {    
    conn := &tcpHAPConnection{
        connection: connection,
        context: context,
    }
    
    // Setup new session for the connection
    session := NewSession(conn)
    context.SetSessionForConnection(session, conn)
    
    return conn
}

func (con *tcpHAPConnection) EncryptedWrite(b []byte) (int, error) {
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

func (con *tcpHAPConnection) DecryptedRead(b []byte) (int, error) {
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

func (con *tcpHAPConnection) Write(b []byte) (int, error) {    
    if con.getEncrypter() != nil {
        return con.EncryptedWrite(b)
    }
    
    return con.connection.Write(b)
}

func (con *tcpHAPConnection) Read(b []byte) (int, error) {
    if con.getDecrypter() != nil {
        return con.DecryptedRead(b)
    }
    
    return con.connection.Read(b)
}

func (con *tcpHAPConnection) Close() error {
    log.Println("[INFO] Close connection and remove session")
    
    // Remove session from the context
    con.context.DeleteSessionForConnection(con.connection)
    
    return con.connection.Close()
}

func (con *tcpHAPConnection) LocalAddr() net.Addr {
    return con.connection.LocalAddr()
}

func (con *tcpHAPConnection) RemoteAddr() net.Addr {
    return con.connection.RemoteAddr()
}

func (con *tcpHAPConnection) SetDeadline(t time.Time) error {
    return con.connection.SetReadDeadline(t)
}

func (con *tcpHAPConnection) SetReadDeadline(t time.Time) error {
    return con.connection.SetReadDeadline(t)
}

func (con *tcpHAPConnection) SetWriteDeadline(t time.Time) error {
    return con.connection.SetWriteDeadline(t)
}

// Helper
func (c *tcpHAPConnection) getEncrypter() Encrypter {
    session  := c.context.GetSessionForConnection(c.connection)
    if session != nil {
        return session.Encrypter()
    }
    
    return nil
}

func (c *tcpHAPConnection) getDecrypter() Decrypter {
    session := c.context.GetSessionForConnection(c.connection)
    if session != nil {
        return session.Decrypter()
    }
    
    return nil
}