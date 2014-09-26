package hapserver

import(
    "net"
    "fmt"
    "time"
    "bytes"
    "io"
    "io/ioutil"
    "bufio"
    "github.com/brutella/hap"
)

type tcpSecureConnection struct {
    connection net.Conn
    context *hap.Context
    
    incomingSecured bool
    decryptedBuffer io.Reader
}

func (con *tcpSecureConnection) SecureWrite(b []byte) (n int, err error) {
    var writeBuffer bytes.Buffer
    writeBuffer.Write(b)
    
    encrypted, err := con.context.SecSession.Encrypt(&writeBuffer)
    if err != nil {
        fmt.Println("Encryption failed", err)
        return 0, err
    }
    
    encrypted_bytes, _ := ioutil.ReadAll(encrypted)
    return con.connection.Write(encrypted_bytes)
}

func (con *tcpSecureConnection) SecureRead(b []byte) (n int, err error) {
    if con.decryptedBuffer == nil {
        sec_conn := bufio.NewReader(con.connection)
        decrypted, err := con.context.SecSession.Decrypt(sec_conn)
        if err != nil {
            fmt.Println("Decryption failed", err)
            return 0, err
        }
        
        con.decryptedBuffer = decrypted
    }
    
    n, err = con.decryptedBuffer.Read(b)
        
    if n < len(b) || err == io.EOF {
        con.decryptedBuffer = nil
    }
    
    return n, err
}

func (con *tcpSecureConnection) Write(b []byte) (n int, err error) {
    // Only encrypt outgoing data when incoming data was secured too
    if con.context.SecSession != nil && con.incomingSecured == true {
        return con.SecureWrite(b)
    }
    
    return con.connection.Write(b)
}

func (con *tcpSecureConnection) Read(b []byte) (n int, err error) {
    if con.context.SecSession != nil || con.incomingSecured == true {
        // Ecrypt everything from now on
        con.incomingSecured = true
        return con.SecureRead(b)
    }
    
    return con.connection.Read(b)
}

func (con *tcpSecureConnection) Close() error {
    return con.connection.Close()
}

func (con *tcpSecureConnection) LocalAddr() net.Addr {
    return con.connection.LocalAddr()
}

func (con *tcpSecureConnection) RemoteAddr() net.Addr {
    return con.connection.RemoteAddr()
}

func (con *tcpSecureConnection) SetDeadline(t time.Time) error {
    return con.connection.SetReadDeadline(t)
}

func (con *tcpSecureConnection) SetReadDeadline(t time.Time) error {
    return con.connection.SetReadDeadline(t)
}

func (con *tcpSecureConnection) SetWriteDeadline(t time.Time) error {
    return con.connection.SetWriteDeadline(t)
}

type TCPHAPListener struct {
    *net.TCPListener
    context *hap.Context
}

func NewTCPHAPListener(l *net.TCPListener, context *hap.Context) *TCPHAPListener {
    return &TCPHAPListener{l, context}
}

func (l *TCPHAPListener) Accept() (c net.Conn, err error) {
    con, err := l.AcceptTCP()
    if err != nil {
        return
    }
    
    return &tcpSecureConnection{connection: con, context: l.context}, nil
}