package gohap

type Client struct {
    name string
    publicKey []byte
}

func NewClient(name string, publicKey []byte) *Client {
    return &Client{name: name, publicKey: publicKey}
}