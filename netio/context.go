package netio

import(
    "net"
    "net/http"
    "fmt"
)

// Provides variables in global has accessible via a connection or request
//
// The value returned by GetConnectionKey() is the same as GetKey()
// if the request comes from the same connection
type Context interface {
    GetKey(c net.Conn) interface{}
    GetConnectionKey(r *http.Request) interface{}
    
    Set(key, val interface{})
    Get(key interface{}) (interface{})
    Delete(key interface{})
}

// Sits on top of a normal context and provides convenient methods to access
// a session for a connection/request
type HAPContext interface {
    Context
    
    // Setter and getter for session
    SetSessionForConnection(s Session, c net.Conn)
    GetSessionForConnection(c net.Conn) Session
    GetSessionForRequest(r *http.Request) Session
    DeleteSessionForConnection(c net.Conn)
    
    // Setter and getter for global bridge
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
    fmt.Println("storage = ", ctx.storage)
}

func (ctx *context) Get(key interface{}) (interface{}) {
    return ctx.storage[key]
}

func (ctx *context) Delete(key interface{}){
    delete(ctx.storage, key)
    fmt.Println("storage = ", ctx.storage)
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

func (ctx *context) SetBridge(b *Bridge) {
    ctx.Set("bridge", b)
}

func (ctx *context) GetBridge() *Bridge {
    return ctx.Get("bridge").(*Bridge)
}