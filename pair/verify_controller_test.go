package pair

import (
    "github.com/brutella/hap"
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
    "bytes"
)

func TestPairVerifyIntegration(t *testing.T) {
    accessory, err := hap.NewAccessory("HAP Test", "123-45-678")
    assert.Nil(t, err)
    
    storage, err := hap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := hap.NewContext(storage)
    controller, err := NewVerifyController(context, accessory)
    assert.Nil(t, err)
    
    client := NewHAPPairVerifyClient("Unit Test", accessory.Password)
    // Setup LTPK for client
    context.SaveClient(hap.NewClient(client.Name,client.PublicKey))
    
    tlvVerifyStart := hap.TLV8Container{}
    tlvVerifyStart.SetByte(hap.TLVType_AuthMethod, 0)
    tlvVerifyStart.SetByte(hap.TLVType_SequenceNumber, VerifyStartRequest)
    tlvVerifyStart.SetBytes(hap.TLVType_PublicKey, client.Session.PublicKey())
    
    reader, err := controller.Handle(tlvVerifyStart.BytesBuffer())
    assert.Nil(t, err)
    
    // Server -> Client
    // - public key
    // - encrypted message
    tlvVerifyResponse, err := hap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, tlvVerifyResponse.GetByte(hap.TLVType_ErrorCode), byte(hap.TLVStatus_NoError))
    assert.Equal(t, tlvVerifyResponse.GetByte(hap.TLVType_SequenceNumber), byte(VerifyStartRespond))
    
    serverPublicKey := tlvVerifyResponse.GetBytes(hap.TLVType_PublicKey)
    assert.NotNil(t, serverPublicKey)
    client.GenerateSharedSecret(serverPublicKey)
    assert.NotNil(t, client.Session.sharedKey)
    assert.NotNil(t, client.Session.encryptionKey)
    
    // Decrypt
    data := tlvVerifyResponse.GetBytes(hap.TLVType_EncryptedData)
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)    
    
    decrypted, err := hap.Chacha20DecryptAndPoly1305Verify(client.Session.EncryptionKey(), []byte("PV-Msg02"), message, mac, nil)
    assert.Nil(t, err)
    
    decrypted_buffer := bytes.NewBuffer(decrypted)
    tlv_decrypted, err := hap.ReadTLV8(decrypted_buffer)
    assert.Nil(t, err)
    
    username  := tlv_decrypted.GetString(hap.TLVType_Username)
    assert.Equal(t, username, "HAP Test")
    // Validate signature
    signature := tlv_decrypted.GetBytes(hap.TLVType_Ed25519Signature)
    
    material := make([]byte, 0)
    material = append(material, serverPublicKey[:]...)
    material = append(material, username...)
    material = append(material, client.Session.PublicKey()...)
    assert.True(t, hap.ValidateED25519Signature(accessory.PublicKey, material, signature))
    
    // Client -> Server
    // encrypted tlv: username and signature
    tlvVerifyFinish := hap.TLV8Container{}
    tlvVerifyFinish.SetByte(hap.TLVType_AuthMethod, 0)
    tlvVerifyFinish.SetByte(hap.TLVType_SequenceNumber, VerifyFinishRequest)
    
    tlv_encrypt := hap.TLV8Container{}
    tlv_encrypt.SetString(hap.TLVType_Username, client.Name)
    
    material = make([]byte, 0)
    material = append(material, client.Session.PublicKey()...)
    material = append(material, []byte(client.Name)...)
    material = append(material, serverPublicKey...)
    
    signature, err = hap.ED25519Signature(client.SecretKey, material)
    assert.Nil(t, err)
    tlv_encrypt.SetBytes(hap.TLVType_Ed25519Signature, signature)
    
    encrypted, mac, _ := hap.Chacha20EncryptAndPoly1305Seal(client.Session.EncryptionKey(), []byte("PV-Msg03"), tlv_encrypt.BytesBuffer().Bytes(), nil)
    
    tlvVerifyFinish.SetBytes(hap.TLVType_EncryptedData, append(encrypted, mac[:]...))
    
    reader, err = controller.Handle(tlvVerifyFinish.BytesBuffer())
    tlvFinishResponse, err := hap.ReadTLV8(reader)
    assert.Nil(t, err)
    assert.Equal(t, tlvFinishResponse.GetByte(hap.TLVType_ErrorCode), byte(hap.TLVStatus_NoError))
    assert.Equal(t, tlvFinishResponse.GetByte(hap.TLVType_SequenceNumber), byte(VerifyFinishRespond))
} 