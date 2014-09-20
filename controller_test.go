package gohap

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "log"
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
)


func SRPClient(username []byte, password []byte) *srp.ClientSession {
    rp, err := srp.NewSRP("openssl.3072", sha512.New, nil)
    cs := rp.NewClientSession(username, password)
    _, _, err = rp.ComputeVerifier(password)
    if err != nil {
        log.Fatal(err)
    }
    
    return cs
}

func pairStartRequestTLV() TLV8Container{
    tlv := TLV8Container{}
    tlv.SetByte(TLVType_AuthMethod, 0)
    tlv.SetByte(TLVType_SequenceNumber, SequenceStartRequest)
    
    return tlv
}

func TestPairing(t *testing.T) {
    controller, err := NewPairingController("Pair-Setup", "123-45-678")
    assert.Nil(t, err)
    
    tlvPairStart := pairStartRequestTLV()
    reader, err := controller.Handle(tlvPairStart.BytesBuffer())
    assert.Nil(t, err)
    
    result, err := ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, result.GetByte(TLVType_ErrorCode), byte(TLVStatus_NoError))
    assert.Equal(t, result.GetByte(TLVType_SequenceNumber), byte(SequenceStartRespond))
    salt := result.GetBytes(TLVType_Salt)
    assert.Equal(t, len(salt), 16) // must be 16 bytes long
    publicKey := result.GetBytes(TLVType_PublicKey)
    assert.NotNil(t, publicKey)
    
    // Simulate client
    // Receive salt `s` and public key `B`
    cs := SRPClient([]byte("Pair-Setup"), []byte("123-45-678"))
    _, err = cs.ComputeKey(salt, publicKey)
    assert.Nil(t, err)
    
    // Generate public key `A` and proof `M1`
    clientPublicKey := cs.GetA()
    clientProof := cs.ComputeAuthenticator() // M1
    
    tlvPairVerify := TLV8Container{}
    tlvPairVerify.SetByte(TLVType_AuthMethod, 0)
    tlvPairVerify.SetByte(TLVType_SequenceNumber, SequenceVerifyRequest)
    tlvPairVerify.SetBytes(TLVType_PublicKey, clientPublicKey)
    tlvPairVerify.SetBytes(TLVType_Proof, clientProof)
    
    // Receive `A` and `M1`
    // Respond `M2`
    reader, err = controller.Handle(tlvPairVerify.BytesBuffer())
    assert.Nil(t, err)
    
    result, err = ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, result.GetByte(TLVType_ErrorCode), byte(TLVStatus_NoError))
    assert.Equal(t, result.GetByte(TLVType_SequenceNumber), byte(SequenceVerifyRespond))
    serverProof := result.GetBytes(TLVType_Proof)
    assert.Equal(t, len(serverProof), len(clientProof))
    
    // Check M2
    assert.True(t, cs.VerifyServerAuthenticator(serverProof))
}  