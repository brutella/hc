package hap

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
func (c *Context) SaveClient(client *Client) {
    c.storage.Set(client.Name + ".ltpk", client.PublicKey)
}