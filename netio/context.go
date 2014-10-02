package netio

import(
    "net"
    "net/http"
    "fmt"
)

type Context interface {
    GetKey(c net.Conn) interface{}
    GetConnectionKey(r *http.Request) interface{}
    
    Set(key, val interface{})
    Get(key interface{}) (interface{})
    Delete(key interface{})
}

type context struct {
    storage map[interface{}]interface{}
}

func NewContext() *context {
    return &context{
        storage: map[interface{}]interface{}{},
    }
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