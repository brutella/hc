package netio

import(
    "net"
)

// TCP listener listens for new connection and creates 
// a new TCP HAP connection for new connections
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
    
    return hapConn, err
}