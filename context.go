package hap

import(
    "io"
)

type SecureSession interface {
    Encrypt(r io.Reader) (io.Reader, error);
    Decrypt(r io.Reader) (io.Reader, error);
}

type Context struct {
    storage Storage
    
    SecSession SecureSession
}

func NewContext(storage Storage) *Context {
    c := Context{storage: storage}
    
    return &c
}

// Returns the client for a specific name
//
// Loads the ltpk from disk and returns initialized client object
func (c *Context) ClientForName(name string) (*Client) {
    data, err := c.storage.Get(name + ".ltpk")
    
    if len(data) > 0 && err == nil{
        client := NewClient(name, data)
        return client
    }
    
    return nil
}

// Stores the long-term public key of the client as {client-name}.ltpk
func (c *Context) SaveClient(client *Client) {
    c.storage.Set(client.Name + ".ltpk", client.PublicKey)
}

func (c *Context) DeleteClient(client *Client) {
    c.storage.Delete(client.Name + ".ltpk")
}


func (c *Context) PublicKeyForAccessory(b *Bridge) []byte {
    return b.PublicKey
}

func (c *Context) SecretKeyForAccessory(b *Bridge) []byte {
    return b.SecretKey
}

func (c *Context) SecureSessionClosed() {
    c.SetSecureSession(nil)
}
func (c *Context) SetSecureSession(secSession SecureSession) {
    c.SecSession = secSession
}