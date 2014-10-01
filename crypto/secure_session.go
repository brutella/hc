package crypto

import(
    "io"
    "encoding/binary"
    "bytes"
    
    "github.com/brutella/hap/common"
)

type secureSession struct {
    encryptKey [32]byte
    decryptKey [32]byte
    
    encryptCount uint64
    decryptCount uint64
}

func NewSecureSessionFromSharedKey(sharedKey [32]byte) (*secureSession, error) {
    salt := []byte("Control-Salt")
    info_out := []byte("Control-Read-Encryption-Key")
    info_in := []byte("Control-Write-Encryption-Key")
    
    var s = new(secureSession)
    var err error
    s.encryptKey, err = HKDF_SHA512(sharedKey[:], salt, info_out)
    s.encryptCount = 0
    if err != nil {
        return nil, err
    }
    
    s.decryptKey, err = HKDF_SHA512(sharedKey[:], salt, info_in)
    s.decryptCount = 0
    
    return s, err
}

// Only used for tests
func NewSecureClientSessionFromSharedKey(sharedKey [32]byte) (*secureSession, error) {
    salt := []byte("Control-Salt")
    info_out := []byte("Control-Write-Encryption-Key")
    info_in := []byte("Control-Read-Encryption-Key")
    
    var s = new(secureSession)
    var err error
    s.encryptKey, err = HKDF_SHA512(sharedKey[:], salt, info_out)
    s.encryptCount = 0
    if err != nil {
        return nil, err
    }
    
    s.decryptKey, err = HKDF_SHA512(sharedKey[:], salt, info_in)
    s.decryptCount = 0
    
    return s, err
}

// Encrypts the data by splitting it into packets
//  [ length (2 bytes)] [ data ] [ auth (16 bytes)]
func (s *secureSession) Encrypt(r io.Reader) (io.Reader, error){
    packets := PacketsFromBytes(r)
    var b bytes.Buffer
    for _, p :=  range packets {
        var nonce_bytes [8]byte
        binary.PutUvarint(nonce_bytes[:], s.encryptCount)
        s.encryptCount += 1
        
        length_bytes := make([]byte, 2)
        binary.LittleEndian.PutUint16(length_bytes, uint16(p.length))
        
        encrypted, mac, err := Chacha20EncryptAndPoly1305Seal(s.encryptKey[:], nonce_bytes[:], p.value, length_bytes[:])
        if err != nil {
            return nil, err
        }
        
        b.Write(length_bytes[:])
        b.Write(encrypted)
        b.Write(mac[:])
    }
    
    return &b, nil
}

// Decrypts the whole thing again
func (s *secureSession) Decrypt(r io.Reader) (io.Reader, error){
    var b bytes.Buffer
    for {
        var length uint16
        if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
        
        if length > 1024 {
            return nil, common.NewErrorf("Packet size too big %d", length)
        }
        
        var buffer = make([]byte, length)
        if err := binary.Read(r, binary.LittleEndian, &buffer); err != nil {
            return nil, err
        }
        
        var mac [16]byte
        if err := binary.Read(r, binary.LittleEndian, &mac); err != nil {
            return nil, err
        }
        
        var nonce_bytes [8]byte
        binary.PutUvarint(nonce_bytes[:], s.decryptCount)
        s.decryptCount += 1
        
        length_bytes := make([]byte, 2)
        binary.LittleEndian.PutUint16(length_bytes, uint16(length))
        
        decrypted, err := Chacha20DecryptAndPoly1305Verify(s.decryptKey[:], nonce_bytes[:], buffer, mac, length_bytes)
        
        if err != nil {
            return nil, common.NewErrorf("Data encryption failed %s", err)
        }
        
        b.Write(decrypted)
        
        if length < PacketLengthMax {
            break
        }
    }
    
    return &b, nil
}