package gohap

type Accessory struct {
    name string
    password string
    publicKey []byte
    secretKey []byte
}

func NewAccessory(name string, password string) (*Accessory, error){
    a := Accessory{name: name, password: password}
    public, secret, err := ED25519GenerateKey(name + password)
    a.publicKey = public
    a.secretKey = secret
    
    return &a, err
}