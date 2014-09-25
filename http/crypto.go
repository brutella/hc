package http

import(
    "io"
    "github.com/brutella/hap"
    "encoding/binary"
    "bytes"
)
// Encrypts the data by splitting it into packets
//  [ length (2 bytes)] [ data ] [ auth (16 bytes)]
func Encrypt(r io.Reader, context *hap.Context) (io.Reader, error){
    packets := PacketsFromBytes(r)
    var b bytes.Buffer
    for _, p :=  range packets {
        var nonce_bytes [8]byte
        binary.PutUvarint(nonce_bytes[:], context.OutCount)
        context.OutCount += 1
        
        var length_bytes [2]byte
        binary.PutUvarint(length_bytes[:], uint64(p.length))
        
        encrypted, mac, err := hap.Chacha20EncryptAndPoly1305Seal(context.OutEncryptionKey[:], nonce_bytes[:], p.value, length_bytes[:])
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
func Decrypt(r io.Reader, context *hap.Context) (io.Reader, error){
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
            return nil, hap.NewErrorf("Packet size too big %d", length)
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
        binary.PutUvarint(nonce_bytes[:], context.OutCount)
        context.OutCount += 1
            
        decrypted, err := hap.Chacha20DecryptAndPoly1305Verify(context.InEncryptionKey[:], nonce_bytes[:], buffer, mac, nil)
        
        if err != nil {
            return nil, hap.NewError("Data encryption failed")
        }
        b.Write(decrypted)
    }
    
    return &b, nil
}