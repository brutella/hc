package db

type Client interface {
    SetName(name string)
    Name() string
    
    SetPublicKey(publicKey []byte)
    PublicKey() []byte
}

type client struct {
    name string
    publicKey []byte
}

func NewClient(name string, publicKey []byte) *client {
    return &client{name: name, publicKey: publicKey}
}

func (c *client) SetName(name string) {
    c.name = name
}

func (c *client) Name() string {
    return c.name
}
    
func (c *client) SetPublicKey(publicKey []byte) {
    c.publicKey = publicKey
}

func (c *client) PublicKey() []byte {
    return c.publicKey
}