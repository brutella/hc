package hap

type Bridge struct {
    Name string
    Password string
    
    PublicKey []byte
    SecretKey []byte
}

func NewBridge(name string, password string) (*Bridge, error){
    b := Bridge{Name: name, Password: password}
    public, secret, err := ED25519GenerateKey(name + password)
    b.PublicKey = public
    b.SecretKey = secret
    
    return &b, err
}