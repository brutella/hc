package netio

import (
	"net"
	"net/http"
	"sync"
)

// Context provides a key-value in-memory storage
type Context interface {
	Set(key, val interface{})
	Get(key interface{}) interface{}
	Delete(key interface{})
}

// HAPContext sits on top of a normal context and provides convenient
// methods to access a session for a specific connection/request
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
	ActiveConnections() []net.Conn

	// Setter and getter for bridge
	SetBridge(b *Bridge)
	GetBridge() *Bridge
}

// HAPContext implementation
type context struct {
	storage map[interface{}]interface{}

	// synchronize access because object is used by different goroutines
	mutex *sync.Mutex
}

func NewContextForBridge(b *Bridge) *context {
	ctx := context{
		storage: map[interface{}]interface{}{},
		mutex:   &sync.Mutex{},
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
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	ctx.storage[key] = val
}

func (ctx *context) Get(key interface{}) interface{} {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	return ctx.storage[key]
}

func (ctx *context) Delete(key interface{}) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	delete(ctx.storage, key)
}

// HAP Context
func (ctx *context) SetSessionForConnection(s Session, c net.Conn) {
	key := ctx.GetKey(c)
	ctx.Set(key, s)
}

func (ctx *context) GetSessionForConnection(c net.Conn) Session {
	key := ctx.GetKey(c)
	if session, ok := ctx.Get(key).(Session); ok == true {
		return session
	}
	return nil
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
func (ctx *context) ActiveConnections() []net.Conn {
	connections := make([]net.Conn, 0)
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
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
