package netio

import(
    "net"
    "os"
    "os/signal"
)

// TCP listener listens for new connection and creates a new HAP connection on new connections
type TCPHAPListener struct {
    *net.TCPListener
    context HAPContext
}

func NewTCPHAPListener(l *net.TCPListener, context HAPContext) *TCPHAPListener {
    return &TCPHAPListener{l, context}
}

func (l *TCPHAPListener) Accept() (c net.Conn, err error) {
    connection, err := l.AcceptTCP()
    if err != nil {
        return
    }
    
    hapConn := NewHAPConnection(connection, l.context)
    closeConnectionOnExit(hapConn)
    
    return hapConn, err
}

func closeConnectionOnExit(connection *tcpHAPConnection) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
            connection.Close()
			os.Exit(0)
		}
	}()
}