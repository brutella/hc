package gohap

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
    "bytes"
)

func TestPairVerifyIntegration(t *testing.T) {
    accessory, err := NewAccessory("HAP Test", "123-45-678")
    assert.Nil(t, err)
    
    storage, err := NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := NewContext(storage)
    controller, err := NewVerifyController(context, accessory)
    assert.Nil(t, err)
    
    client := NewHAPPairVerifyClient("Unit Test", accessory.password)
    // Setup LTPK for client
    context.SaveClient(&Client{name:client.name, publicKey: client.publicKey})
    
    tlvVerifyStart := TLV8Container{}
    tlvVerifyStart.SetByte(TLVType_AuthMethod, 0)
    tlvVerifyStart.SetByte(TLVType_SequenceNumber, VerifyStartRequest)
    tlvVerifyStart.SetBytes(TLVType_PublicKey, client.session.publicKey[:])
    
    reader, err := controller.Handle(tlvVerifyStart.BytesBuffer())
    assert.Nil(t, err)
    
    // Server -> Client
    // - public key
    // - encrypted message
    tlvVerifyResponse, err := ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, tlvVerifyResponse.GetByte(TLVType_ErrorCode), byte(TLVStatus_NoError))
    assert.Equal(t, tlvVerifyResponse.GetByte(TLVType_SequenceNumber), byte(VerifyStartRespond))
    
    serverPublicKey := tlvVerifyResponse.GetBytes(TLVType_PublicKey)
    assert.NotNil(t, serverPublicKey)
    client.GenerateSharedSecret(serverPublicKey)
    assert.NotNil(t, client.session.sharedKey)
    assert.NotNil(t, client.session.encryptionKey)
    
    // Decrypt
    data := tlvVerifyResponse.GetBytes(TLVType_EncryptedData)
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)    
    
    decrypted, err := Chacha20DecryptAndPoly1305Verify(client.session.encryptionKey[:], []byte("PV-Msg02"), message, mac, nil)
    assert.Nil(t, err)
    
    decrypted_buffer := bytes.NewBuffer(decrypted)
    tlv_decrypted, err := ReadTLV8(decrypted_buffer)
    assert.Nil(t, err)
    
    username  := tlv_decrypted.GetString(TLVType_Username)
    assert.Equal(t, username, "HAP Test")
    // Validate signature
    signature := tlv_decrypted.GetBytes(TLVType_Ed25519Signature)
    
    material := make([]byte, 0)
    material = append(material, serverPublicKey[:]...)
    material = append(material, username...)
    material = append(material, client.session.publicKey[:]...)
    assert.True(t, ValidateED25519Signature(accessory.publicKey, material, signature))
    
    // Client -> Server
    // encrypted tlv: username and signature
    tlvVerifyFinish := TLV8Container{}
    tlvVerifyFinish.SetByte(TLVType_AuthMethod, 0)
    tlvVerifyFinish.SetByte(TLVType_SequenceNumber, VerifyFinishRequest)
    
    tlv_encrypt := TLV8Container{}
    tlv_encrypt.SetString(TLVType_Username, client.name)
    
    material = make([]byte, 0)
    material = append(material, client.session.publicKey[:]...)
    material = append(material, serverPublicKey...)
    
    signature, err = ED25519Signature(client.secretKey, material)
    assert.Nil(t, err)
    tlv_encrypt.SetBytes(TLVType_Ed25519Signature, signature)
    
    encrypted, mac, _ := Chacha20EncryptAndPoly1305Seal(client.session.encryptionKey[:], []byte("PV-Msg03"), tlv_encrypt.BytesBuffer().Bytes(), mac, nil)
    
    tlvVerifyFinish.SetBytes(TLVType_EncryptedData, append(encrypted, mac[:]...))
    
    reader, err = controller.Handle(tlvVerifyFinish.BytesBuffer())
    tlvFinishResponse, err := ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, tlvFinishResponse.GetByte(TLVType_ErrorCode), byte(TLVStatus_NoError))
    assert.Equal(t, tlvFinishResponse.GetByte(TLVType_SequenceNumber), byte(VerifyFinishRespond))
} 