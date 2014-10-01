package hap

import(
    "github.com/brutella/hap/common"
)

type Context struct {
    storage Storage
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
func (c *Context) SaveClient(client *Client) error {
    if len(client.PublicKey) == 0 {
        return common.NewErrorf("No public key to save for client%s\n", client.Name)
    }
    
    return c.storage.Set(client.Name + ".ltpk", client.PublicKey)
}

func (c *Context) DeleteClient(client *Client) {
    c.storage.Delete(client.Name + ".ltpk")
}