package hap

type Context struct {
    storage Storage
}

func NewContext(storage Storage) *Context {
    c := Context{storage: storage}
    
    return &c
}

func (c *Context) ClientForName(name string) (*Client) {
    data, err := c.storage.Get(name)
    
    if len(data) > 0 && err == nil{
        client := NewClient(name, data)
        return client
    }
    
    return nil
}

func (c *Context) SaveClient(client *Client) {
    c.storage.Set(client.Name, client.PublicKey)
}