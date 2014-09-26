package server

import(
    "net"
    "fmt"
    "time"
    "bytes"
    "io"
    "net/http"
    "io/ioutil"
    "bufio"
    "github.com/brutella/hap"
)

// Creates a server which handles tcp connection based on HAP protocol
// The tcp connection is secured after the crypto keys are verified
func ListenAndServe(addr string, handler http.Handler, context *hap.Context) error {
    server := http.Server{Addr: addr, Handler:handler}
    ln, err := net.Listen("tcp", server.Addr)
    if err != nil {
        return err
    }
    
    listener := NewTCPHAPListener(ln.(*net.TCPListener), context)
    
    return server.Serve(listener)
}

// TCP connection based on HAP protocol
// The connection in secured after a secured session is set up
type tcpHAPConnection struct {
    connection net.Conn
    context *hap.Context
    
    isEncrypted bool
    decryptedBuffer io.Reader
}

// Encrypts outgoing data before writing to connection
func (con *tcpHAPConnection) SecureWrite(b []byte) (n int, err error) {
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

// Decrypts incoming data
func (con *tcpHAPConnection) SecureRead(b []byte) (n int, err error) {
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

func (con *tcpHAPConnection) Write(b []byte) (n int, err error) {
    // Only encrypt outgoing data when incoming data was secured too
    if con.isEncrypted == true {
        if con.context.SecSession == nil {
            err := hap.NewError("[ERROR] Can not write to secured connection without crypto keys")
            fmt.Println(err)
            return 0, err
        }
        return con.SecureWrite(b)
    }
    
    return con.connection.Write(b)
}

func (con *tcpHAPConnection) Read(b []byte) (n int, err error) {
    if con.context.SecSession != nil {
        // Decrypt request and ecrypt responds
        con.isEncrypted = true
        return con.SecureRead(b)
    }
    
    con.isEncrypted = false
    
    return con.connection.Read(b)
}

func (con *tcpHAPConnection) Close() error {
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

// TCP listener which creates a tcp HAP connection which uses a secured session
// to communicate securily
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
    
    return &tcpHAPConnection{connection: con, context: l.context}, nil
}