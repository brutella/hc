package pair

import (
    "github.com/brutella/hap"
    
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
    "fmt"
    "encoding/hex"
)

func TestPairingIntegration(t *testing.T) {
    accessory, err := hap.NewAccessory("HAP Test", "123-45-678")
    assert.Nil(t, err)
    
    storage, err := hap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := hap.NewContext(storage)
    controller, err := NewSetupController(context, accessory)
    assert.Nil(t, err)
    
    tlvPairStart := hap.TLV8Container{}
    tlvPairStart.SetByte(hap.TLVType_AuthMethod, 0)
    tlvPairStart.SetByte(hap.TLVType_SequenceNumber, SequenceStartRequest)
    
    reader, err := controller.Handle(tlvPairStart.BytesBuffer())
    assert.Nil(t, err)
    
    result, err := hap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, result.GetByte(hap.TLVType_ErrorCode), byte(hap.TLVStatus_NoError))
    assert.Equal(t, result.GetByte(hap.TLVType_SequenceNumber), byte(SequenceStartRespond))
    salt := result.GetBytes(hap.TLVType_Salt)
    assert.Equal(t, len(salt), 16) // 16 bytes
    publicKey := result.GetBytes(hap.TLVType_PublicKey)
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
    
    tlvPairVerify := hap.TLV8Container{}
    tlvPairVerify.SetByte(hap.TLVType_AuthMethod, 0)
    tlvPairVerify.SetByte(hap.TLVType_SequenceNumber, SequenceVerifyRequest)
    tlvPairVerify.SetBytes(hap.TLVType_PublicKey, clientPublicKey)
    tlvPairVerify.SetBytes(hap.TLVType_Proof, clientProof)
    
    // Server
    // 1) Receive `A` and `M1`
    // 2) Send `M2`
    reader, err = controller.Handle(tlvPairVerify.BytesBuffer())
    assert.Nil(t, err)
    
    result, err = hap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, result.GetByte(hap.TLVType_ErrorCode), byte(hap.TLVStatus_NoError))
    assert.Equal(t, result.GetByte(hap.TLVType_SequenceNumber), byte(SequenceVerifyRespond))
    serverProof := result.GetBytes(hap.TLVType_Proof)
    assert.Equal(t, len(serverProof), len(clientProof))
    
    // Client
    // 1) Check M2
    assert.True(t, client.Session.VerifyServerAuthenticator(serverProof))
    
    // 2) Send username, LTPK, proof as encrypted message
    H2, err := hap.HKDF_SHA512(clientSecretKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
    material := make([]byte, 0)
    material = append(material, H2[:]...)
    material = append(material, client.Name...)
    material = append(material, client.PublicKey...)
    
    signature, err := hap.ED25519Signature(client.SecretKey, material)
    assert.Nil(t, err)
    tlvPairKeyExchange := hap.TLV8Container{}
    tlvPairKeyExchange.SetString(hap.TLVType_Username, client.Name)
    tlvPairKeyExchange.SetBytes(hap.TLVType_PublicKey, []byte(client.PublicKey))
    tlvPairKeyExchange.SetBytes(hap.TLVType_Ed25519Signature, []byte(signature))
    
    K, err := hap.HKDF_SHA512(clientSecretKey, []byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
    assert.Nil(t, err)
    fmt.Println("K:", hex.EncodeToString(K[:]))
    
    encrypted, tag, err := hap.Chacha20EncryptAndPoly1305Seal(K[:], []byte("PS-Msg05"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)
    assert.Nil(t, err)
    
    tlvKeyExchangeRequest := hap.TLV8Container{}
    tlvKeyExchangeRequest.SetByte(hap.TLVType_AuthMethod, 0)
    tlvKeyExchangeRequest.SetByte(hap.TLVType_SequenceNumber, SequenceKeyExchangeRequest)
    tlvKeyExchangeRequest.SetBytes(hap.TLVType_EncryptedData, append(encrypted, tag[:]...))
    
    // Server response with
    // - Encrypted message
    reader, err = controller.Handle(tlvKeyExchangeRequest.BytesBuffer())
    assert.Nil(t, err)
    
    keyVerifyResponse, err := hap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, keyVerifyResponse.GetByte(hap.TLVType_ErrorCode), byte(hap.TLVStatus_NoError))
    assert.Equal(t, keyVerifyResponse.GetByte(hap.TLVType_SequenceNumber), byte(SequenceKeyExchangeRepond))
    keyVerifyResponseEncrypted := keyVerifyResponse.GetBytes(hap.TLVType_EncryptedData)
    assert.NotNil(t, keyVerifyResponseEncrypted)
    
    // TODO verify response, encrpyted data, signature,...
}  