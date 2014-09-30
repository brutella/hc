package hap

type Bridge struct {
    name string
    password string
    info BridgeInfo
    
    PublicKey []byte
    SecretKey []byte
}

// Creates a new bridge based on the provided info
// The long-term public and secret key are based on the serial number which should be unique for every bridge
func NewBridge(info BridgeInfo) (*Bridge, error){
    b := Bridge{info: info}
    public, secret, err := ED25519GenerateKey(b.info.SerialNumber)
    b.PublicKey = public
    b.SecretKey = secret
    
    return &b, err
}

func (b *Bridge) Name() string {
    return b.info.Name
}

func (b *Bridge) Id() string {
    return b.info.Id
}

func (b *Bridge) Password() string {
    return b.info.Password
}