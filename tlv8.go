package gohap

import(
    "encoding/binary"
    "bytes"
    "io"
)

type TLV8Container struct {
    Items []tlv8
}

func ReadTLV8(r io.Reader) (*TLV8Container, error) {
    var items = make([]tlv8, 0, 1)
    for {
        var item tlv8
        if err := binary.Read(r, binary.LittleEndian, &item.tag); err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
        if err := binary.Read(r, binary.LittleEndian, &item.length); err != nil {
            return nil, err
        }
        
        item.value = make([]byte, item.length)
        if _, err := io.ReadFull(r, item.value); err != nil {
            return nil, err
        }
        
        items = append(items, item)
    }
    
    return &TLV8Container{Items:items}, nil
}

func (t *TLV8Container) GetBuffer(tag uint8) *bytes.Buffer {
    // TODO append
    var b bytes.Buffer
    for _, item := range t.Items {
        if item.tag == tag {
            b.Write(item.value)
        }
    }
    
    return &b
}

func (t *TLV8Container) GetBytes(tag uint8) []byte {
    return t.GetBuffer(tag).Bytes()
}

func (t *TLV8Container) GetUInt64(tag uint8) uint64 {
    integer, _ :=  binary.ReadUvarint(t.GetBuffer(tag))
    
    return integer
}

func (t *TLV8Container) Set(tag uint8, value []byte) {
    r := bytes.NewBuffer(value)
    
    for {
        var item = tlv8{}
        item.tag = tag
        var bytes = make([]byte, 255)
        n, err := io.ReadFull(r, bytes)
        if err == nil || err == io.ErrUnexpectedEOF {
            item.length = uint8(n)
            item.value = bytes[:item.length]
            t.Items = append(t.Items, item)
            
            if err == io.ErrUnexpectedEOF { // Fewer than 255 bytes read
                break
            }
        } else {
            break
        }
    }
}

func (t *TLV8Container) BytesBuffer() *bytes.Buffer {
    var b bytes.Buffer
    for _, item := range t.Items {
        b.Write([]byte{item.tag})
        b.Write([]byte{item.length})
        b.Write(item.value)
    }
    
    return &b
}

// Encodes data into by tag, length and value
type tlv8 struct {
    tag uint8
    length uint8
    value []byte 
}