package gohap

type PairVerifySession struct {
    clientPublicKey [32]byte
    publicKey [32]byte
    secretKey [32]byte
    sharedKey [32]byte
    encryptionKey [32]byte
}

func NewPairVerifySession() (*PairVerifySession) {
    return &PairVerifySession{}
}