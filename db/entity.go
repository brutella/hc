package db

// Entity represents a HomeKit entity (e.g. iOS device or HomeKit bridge).
type Entity interface {
	// Name returns the entity name
	Name() string

	// SetName sets the entity name
	SetName(name string)

	// PublicKey returns the entity (long-term) public key
	PublicKey() []byte

	// SetPublicKey sets the entity (long-term) public key
	SetPublicKey(publicKey []byte)
}

type entity struct {
	name      string
	publicKey []byte
}

// NewEntity returns a entity with a name and public key.
func NewEntity(name string, publicKey []byte) Entity {
	return &entity{name: name, publicKey: publicKey}
}

func (c *entity) SetName(name string) {
	c.name = name
}

func (c *entity) Name() string {
	return c.name
}

func (c *entity) SetPublicKey(publicKey []byte) {
	c.publicKey = publicKey
}

func (c *entity) PublicKey() []byte {
	return c.publicKey
}
