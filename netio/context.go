package netio

import(
    "net"
    "net/http"
)

// Provides variables in global has accessible via a connection or request
type Context interface {    
    Set(key, val interface{})
    Get(key interface{}) (interface{})
    Delete(key interface{})
}

// Sits on top of a normal context and provides convenient 
// methods to access a session for a connection/request
type HAPContext interface {
    Context
    
    // Returns a key to uniquely identify the connection
    GetKey(c net.Conn) interface{}
    
    // Returns the same key as for the underlying connection
    GetConnectionKey(r *http.Request) interface{}
    
    // Setter and getter for session
    SetSessionForConnection(s Session, c net.Conn)
    GetSessionForConnection(c net.Conn) Session
    GetSessionForRequest(r *http.Request) Session
    DeleteSessionForConnection(c net.Conn)
    
    // Returns a list of active connections
    ActiveConnection() []net.Conn
    
    // Setter and getter for bridge
    SetBridge(b *Bridge)
    GetBridge() *Bridge
}

// HAPContext implementation
type context struct {
    storage map[interface{}]interface{}
}

func NewContextForBridge(b *Bridge) *context {
    ctx := context{
        storage: map[interface{}]interface{}{},
    }
    ctx.SetBridge(b)
    return &ctx
}

func (ctx *context) GetKey(c net.Conn) interface{} {
    return c.RemoteAddr().String()
}

func (ctx *context) GetConnectionKey(r *http.Request) interface{} {
    return r.RemoteAddr
}

func (ctx *context) Set(key, val interface{}) {
    ctx.storage[key] = val
}

func (ctx *context) Get(key interface{}) (interface{}) {
    return ctx.storage[key]
}

func (ctx *context) Delete(key interface{}){
    delete(ctx.storage, key)
}

// HAP Context
func (ctx *context) SetSessionForConnection(s Session, c net.Conn) {
    key := ctx.GetKey(c)
    ctx.Set(key, s)
}

func (ctx *context) GetSessionForConnection(c net.Conn) Session {
    key := ctx.GetKey(c)
    return ctx.Get(key).(Session)
}

func (ctx *context) GetSessionForRequest(r *http.Request) Session {
    key := ctx.GetConnectionKey(r)
    return ctx.Get(key).(Session)
}

func (ctx *context) DeleteSessionForConnection(c net.Conn) {
    key := ctx.GetKey(c)
    ctx.Delete(key)
}

// Returns a list of active connections
func (ctx *context) ActiveConnection() []net.Conn {
    connections := make([]net.Conn, 0)
    
    for _, v := range ctx.storage {
        if s, ok := v.(Session); ok == true {
            connections = append(connections, s.Connection())
        }
    }
    
    return connections
}

func (ctx *context) SetBridge(b *Bridge) {
    ctx.Set("bridge", b)
}

func (ctx *context) GetBridge() *Bridge {
    return ctx.Get("bridge").(*Bridge)
}