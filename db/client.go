package db

type Client struct {
    Name string
    PublicKey []byte
}

func NewClient(name string, publicKey []byte) *Client {
    return &Client{Name: name, PublicKey: publicKey}
}