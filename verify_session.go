package gohap

type PairVerifySession struct {
    storage Storage
    
    publicKey [32]byte
    secretKey [32]byte
    encryptionKey [32]byte // K
}

func NewPairVerifySession(s Storage) (*PairVerifySession) {
    return &PairVerifySession{storage: s}
}