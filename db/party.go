package db

type Party struct {
    Name string
    SerialNumber string
    PublicKey []byte
    SecretKey []byte
}

func NewParty(name, serial string, publicKey []byte, secretKey []byte) *Party {
    return &Party{Name: name, SerialNumber: serial, PublicKey: publicKey, SecretKey: secretKey}
}