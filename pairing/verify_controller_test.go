package gohap

import (
    "github.com/brutella/gohap"
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
    "bytes"
)

func TestPairVerifyIntegration(t *testing.T) {
    accessory, err := gohap.NewAccessory("HAP Test", "123-45-678")
    assert.Nil(t, err)
    
    storage, err := gohap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := gohap.NewContext(storage)
    controller, err := NewVerifyController(context, accessory)
    assert.Nil(t, err)
    
    client := NewHAPPairVerifyClient("Unit Test", accessory.Password)
    // Setup LTPK for client
    context.SaveClient(gohap.NewClient(client.Name,client.PublicKey))
    
    tlvVerifyStart := gohap.TLV8Container{}
    tlvVerifyStart.SetByte(gohap.TLVType_AuthMethod, 0)
    tlvVerifyStart.SetByte(gohap.TLVType_SequenceNumber, VerifyStartRequest)
    tlvVerifyStart.SetBytes(gohap.TLVType_PublicKey, client.Session.PublicKey())
    
    reader, err := controller.Handle(tlvVerifyStart.BytesBuffer())
    assert.Nil(t, err)
    
    // Server -> Client
    // - public key
    // - encrypted message
    tlvVerifyResponse, err := gohap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, tlvVerifyResponse.GetByte(gohap.TLVType_ErrorCode), byte(gohap.TLVStatus_NoError))
    assert.Equal(t, tlvVerifyResponse.GetByte(gohap.TLVType_SequenceNumber), byte(VerifyStartRespond))
    
    serverPublicKey := tlvVerifyResponse.GetBytes(gohap.TLVType_PublicKey)
    assert.NotNil(t, serverPublicKey)
    client.GenerateSharedSecret(serverPublicKey)
    assert.NotNil(t, client.Session.sharedKey)
    assert.NotNil(t, client.Session.encryptionKey)
    
    // Decrypt
    data := tlvVerifyResponse.GetBytes(gohap.TLVType_EncryptedData)
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)    
    
    decrypted, err := gohap.Chacha20DecryptAndPoly1305Verify(client.Session.EncryptionKey(), []byte("PV-Msg02"), message, mac, nil)
    assert.Nil(t, err)
    
    decrypted_buffer := bytes.NewBuffer(decrypted)
    tlv_decrypted, err := gohap.ReadTLV8(decrypted_buffer)
    assert.Nil(t, err)
    
    username  := tlv_decrypted.GetString(gohap.TLVType_Username)
    assert.Equal(t, username, "HAP Test")
    // Validate signature
    signature := tlv_decrypted.GetBytes(gohap.TLVType_Ed25519Signature)
    
    material := make([]byte, 0)
    material = append(material, serverPublicKey[:]...)
    material = append(material, username...)
    material = append(material, client.Session.PublicKey()...)
    assert.True(t, gohap.ValidateED25519Signature(accessory.PublicKey, material, signature))
    
    // Client -> Server
    // encrypted tlv: username and signature
    tlvVerifyFinish := gohap.TLV8Container{}
    tlvVerifyFinish.SetByte(gohap.TLVType_AuthMethod, 0)
    tlvVerifyFinish.SetByte(gohap.TLVType_SequenceNumber, VerifyFinishRequest)
    
    tlv_encrypt := gohap.TLV8Container{}
    tlv_encrypt.SetString(gohap.TLVType_Username, client.Name)
    
    material = make([]byte, 0)
    material = append(material, client.Session.PublicKey()...)
    material = append(material, []byte(client.Name)...)
    material = append(material, serverPublicKey...)
    
    signature, err = gohap.ED25519Signature(client.SecretKey, material)
    assert.Nil(t, err)
    tlv_encrypt.SetBytes(gohap.TLVType_Ed25519Signature, signature)
    
    encrypted, mac, _ := gohap.Chacha20EncryptAndPoly1305Seal(client.Session.EncryptionKey(), []byte("PV-Msg03"), tlv_encrypt.BytesBuffer().Bytes(), nil)
    
    tlvVerifyFinish.SetBytes(gohap.TLVType_EncryptedData, append(encrypted, mac[:]...))
    
    reader, err = controller.Handle(tlvVerifyFinish.BytesBuffer())
    tlvFinishResponse, err := gohap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, tlvFinishResponse.GetByte(gohap.TLVType_ErrorCode), byte(gohap.TLVStatus_NoError))
    assert.Equal(t, tlvFinishResponse.GetByte(gohap.TLVType_SequenceNumber), byte(VerifyFinishRespond))
} 