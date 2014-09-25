package hap

type Context struct {
    storage Storage
    
    SharedKey [32]byte // established on key-verification phase
    OutEncryptionKey [32]byte // for outgoing data
    OutCount uint64
    
    InEncryptionKey [32]byte // for incoming data
    InCount uint64
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

func (c *Context) PublicKeyForAccessory(b *Bridge) []byte {
    return b.PublicKey
}

func (c *Context) SecretKeyForAccessory(b *Bridge) []byte {
    return b.SecretKey
}

func (c *Context) GenerateEncryptionKeysWithSharedkey(sharedKey [32]byte) error {
    c.SharedKey = sharedKey
    salt := []byte("Control-Salt")
    
    info_in := []byte("Control-Read-Encryption-Key")
    info_out := []byte("Control-Write-Encryption-Key")
    
    var err error
    c.OutEncryptionKey, err = HKDF_SHA512(c.SharedKey[:], salt, info_out)
    c.OutCount = 0
    if err != nil {
        return err
    }
    
    c.InEncryptionKey, err = HKDF_SHA512(c.SharedKey[:], salt, info_in)
    c.InCount = 0
    
    return err
}