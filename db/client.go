package db

// Client represents a HomeKit client (e.g. iOS device).
type Client interface {
	// Name returns the client name
	Name() string

	// SetName sets the client name
	SetName(name string)

	// PublicKey returns the client (long-term) public key
	PublicKey() []byte

	// SetPublicKey sets the client (long-term) public key
	SetPublicKey(publicKey []byte)
}

type client struct {
	name      string
	publicKey []byte
}

// NewClient returns a client with a name and public key.
func NewClient(name string, publicKey []byte) Client {
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
