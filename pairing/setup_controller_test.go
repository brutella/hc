package gohap

import (
    "github.com/brutella/gohap"
    
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
    "fmt"
    "encoding/hex"
)

func TestPairingIntegration(t *testing.T) {
    accessory, err := gohap.NewAccessory("HAP Test", "123-45-678")
    assert.Nil(t, err)
    
    storage, err := gohap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := gohap.NewContext(storage)
    controller, err := NewSetupController(context, accessory)
    assert.Nil(t, err)
    
    tlvPairStart := gohap.TLV8Container{}
    tlvPairStart.SetByte(gohap.TLVType_AuthMethod, 0)
    tlvPairStart.SetByte(gohap.TLVType_SequenceNumber, SequenceStartRequest)
    
    reader, err := controller.Handle(tlvPairStart.BytesBuffer())
    assert.Nil(t, err)
    
    result, err := gohap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, result.GetByte(gohap.TLVType_ErrorCode), byte(gohap.TLVStatus_NoError))
    assert.Equal(t, result.GetByte(gohap.TLVType_SequenceNumber), byte(SequenceStartRespond))
    salt := result.GetBytes(gohap.TLVType_Salt)
    assert.Equal(t, len(salt), 16) // 16 bytes
    publicKey := result.GetBytes(gohap.TLVType_PublicKey)
    assert.Equal(t, len(publicKey), 384) // 384 bytes
    
    // Client
    // 1) Receive salt `s` and public key `B`
    client := NewHAPPairSetupClient("Unit Test", accessory.Password)
    clientSecretKey, err := client.Session.ComputeKey(salt, publicKey)
    assert.Nil(t, err)
    assert.NotNil(t, clientSecretKey)
    
    // 2) Send public key `A` and proof `M1`
    clientPublicKey := client.Session.GetA() // SRP public key
    clientProof := client.Session.ComputeAuthenticator() // M1
    
    tlvPairVerify := gohap.TLV8Container{}
    tlvPairVerify.SetByte(gohap.TLVType_AuthMethod, 0)
    tlvPairVerify.SetByte(gohap.TLVType_SequenceNumber, SequenceVerifyRequest)
    tlvPairVerify.SetBytes(gohap.TLVType_PublicKey, clientPublicKey)
    tlvPairVerify.SetBytes(gohap.TLVType_Proof, clientProof)
    
    // Server
    // 1) Receive `A` and `M1`
    // 2) Send `M2`
    reader, err = controller.Handle(tlvPairVerify.BytesBuffer())
    assert.Nil(t, err)
    
    result, err = gohap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, result.GetByte(gohap.TLVType_ErrorCode), byte(gohap.TLVStatus_NoError))
    assert.Equal(t, result.GetByte(gohap.TLVType_SequenceNumber), byte(SequenceVerifyRespond))
    serverProof := result.GetBytes(gohap.TLVType_Proof)
    assert.Equal(t, len(serverProof), len(clientProof))
    
    // Client
    // 1) Check M2
    assert.True(t, client.Session.VerifyServerAuthenticator(serverProof))
    
    // 2) Send username, LTPK, proof as encrypted message
    H2, err := gohap.HKDF_SHA512(clientSecretKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
    material := make([]byte, 0)
    material = append(material, H2[:]...)
    material = append(material, client.Name...)
    material = append(material, client.PublicKey...)
    
    signature, err := gohap.ED25519Signature(client.SecretKey, material)
    assert.Nil(t, err)
    tlvPairKeyExchange := gohap.TLV8Container{}
    tlvPairKeyExchange.SetString(gohap.TLVType_Username, client.Name)
    tlvPairKeyExchange.SetBytes(gohap.TLVType_PublicKey, []byte(client.PublicKey))
    tlvPairKeyExchange.SetBytes(gohap.TLVType_Ed25519Signature, []byte(signature))
    
    K, err := gohap.HKDF_SHA512(clientSecretKey, []byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
    assert.Nil(t, err)
    fmt.Println("K:", hex.EncodeToString(K[:]))
    
    encrypted, tag, err := gohap.Chacha20EncryptAndPoly1305Seal(K[:], []byte("PS-Msg05"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)
    assert.Nil(t, err)
    
    tlvKeyExchangeRequest := gohap.TLV8Container{}
    tlvKeyExchangeRequest.SetByte(gohap.TLVType_AuthMethod, 0)
    tlvKeyExchangeRequest.SetByte(gohap.TLVType_SequenceNumber, SequenceKeyExchangeRequest)
    tlvKeyExchangeRequest.SetBytes(gohap.TLVType_EncryptedData, append(encrypted, tag[:]...))
    
    // Server response with
    // - Encrypted message
    reader, err = controller.Handle(tlvKeyExchangeRequest.BytesBuffer())
    assert.Nil(t, err)
    
    keyVerifyResponse, err := gohap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, keyVerifyResponse.GetByte(gohap.TLVType_ErrorCode), byte(gohap.TLVStatus_NoError))
    assert.Equal(t, keyVerifyResponse.GetByte(gohap.TLVType_SequenceNumber), byte(SequenceKeyExchangeRepond))
    keyVerifyResponseEncrypted := keyVerifyResponse.GetBytes(gohap.TLVType_EncryptedData)
    assert.NotNil(t, keyVerifyResponseEncrypted)
    
    // TODO verify response, encrpyted data, signature,...
}  