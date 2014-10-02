package netio

import(
    "net"
    "net/http"
)

// Creates a http.Server which to handle TCP connection
// 
// This server uses a TCPHAPListener to accept incoming connections.
//
// Overview over components
//      [server]  --------------------------> [mux]
//          |                                   |
//          | create listener                   |
//          |                                   |
//          V                                   |  (multiplexes to endpoints)
// -> [listener] (accept new connection)        |
//          |                                   |
//          | (create new connection)           |
//          |                                   |
//          V                                   v
// <-> [connection] (read and write)  <---- [endpoint]
//          |                                   |
//          | (access session)                  | (access controller based on connection/request)
//          |                                   |
//          V                                   |
//       context (provides variables) <---------
func ListenAndServe(addr string, handler http.Handler, context HAPContext) error {
    server := http.Server{Addr: addr, Handler:handler}
    ln, err := net.Listen("tcp", server.Addr)
    if err != nil {
        return err
    }
    
    listener := NewTCPHAPListener(ln.(*net.TCPListener), context)
    
    return server.Serve(listener)
}