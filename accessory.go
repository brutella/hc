package hap

type Accessory struct {
    UUId int `json:"aid"`
    Services []Service `json:"services"`
    
    Name string `json:"-"`
    Password string `json:"-"`
    
    publicKey []byte
    secretKey []byte
}

func NewAccessory(name string, password string) (*Accessory, error){
    a := Accessory{Name: name, Password: password}
    public, secret, err := ED25519GenerateKey(name + password)
    a.publicKey = public
    a.secretKey = secret
    
    return &a, err
}