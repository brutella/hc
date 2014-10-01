package pair

const (
    WaitingForRequest  = 0x00
    
    PairStartRequest       = 0x01
    PairStartRespond       = 0x02
    PairVerifyRequest      = 0x03
    PairVerifyRespond      = 0x04
    PairKeyExchangeRequest = 0x05
    PairKeyExchangeRespond = 0x06
    
    VerifyStartRequest  = 0x01
    VerifyStartRespond  = 0x02
    VerifyFinishRequest = 0x03
    VerifyFinishRespond = 0x04
)