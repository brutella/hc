package netio

import(
    "fmt"
    "time"
    "os"
    "os/signal"
    "net"
    "net/http"
    "bytes"
    
    "io/ioutil"
    "io"
    "bufio"
)

// Creates a server which handles tcp connection based on HAP protocol
// The tcp connection is secured after the crypto keys are verified
func ListenAndServe(addr string, handler http.Handler, context Context) error {
    server := http.Server{Addr: addr, Handler:handler}
    ln, err := net.Listen("tcp4", server.Addr)
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
    context Context
    readBuffer io.Reader
}

func (c *tcpHAPConnection) GetEncrypter() Encrypter {
    key     := c.context.GetKey(c.connection)    
    session  := c.context.Get(key).(Session)
    return session.Encrypter()
}

func (c *tcpHAPConnection) GetDecrypter() Decrypter {
    key     := c.context.GetKey(c.connection)    
    session  := c.context.Get(key).(Session)
    return session.Decrypter()
}

func (con *tcpHAPConnection) Write(b []byte) (int, error) {    
    var buffer bytes.Buffer
    buffer.Write(b)
    encrypted, err := con.GetEncrypter().Encrypt(&buffer)
    
    if err != nil {
        err = con.connection.Close()
        return 0, err
    }
    
    encryptedBytes, err := ioutil.ReadAll(encrypted)    
    n, err := con.connection.Write(encryptedBytes)
    
    return n, err
}

func (con *tcpHAPConnection) Read(b []byte) (int, error) {
    if con.readBuffer == nil {
        buffered := bufio.NewReader(con.connection)
        decrypted, err := con.GetDecrypter().Decrypt(buffered)
        if err != nil {
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

func (con *tcpHAPConnection) Close() error {
    fmt.Println("Close connection and remove session")
    
    // Delete the session for the connetion
    key := con.context.GetKey(con.connection)
    con.context.Delete(key)
    
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
    context Context
}

func NewTCPHAPListener(l *net.TCPListener, context Context) *TCPHAPListener {
    return &TCPHAPListener{l, context}
}

func (l *TCPHAPListener) Accept() (c net.Conn, err error) {
    con, err := l.AcceptTCP()
    if err != nil {
        return
    }
    
    session := NewSession()
    key := l.context.GetKey(con)
    l.context.Set(key, session)
    
    fmt.Println("***** New session")
    
    hapConn, err := &tcpHAPConnection{connection: con, context: l.context}, nil
    if err == nil {
        closeConnectionOnExit(hapConn)
    }
    
    return hapConn, err
}

func closeConnectionOnExit(connection *tcpHAPConnection) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Println("Closing...")
            connection.Close()
			os.Exit(0)
		}
	}()
}