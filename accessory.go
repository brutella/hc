package hap

type Accessory struct {
    Name string
    Password string
    PublicKey []byte
    SecretKey []byte
}

func NewAccessory(name string, password string) (*Accessory, error){
    a := Accessory{Name: name, Password: password}
    public, secret, err := ED25519GenerateKey(name + password)
    a.PublicKey = public
    a.SecretKey = secret
    
    return &a, err
}