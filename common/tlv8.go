package common

import(
    "encoding/binary"
    "bytes"
    "io"
)

type tlv8Container struct {
    Container
    Items []tlv8
}
func NewTLV8Container() Container {
    return &tlv8Container{
        Items:make([]tlv8, 0, 1),
    }
}

func NewTLV8ContainerFromReader(r io.Reader) (Container, error) {
    var items = make([]tlv8, 0, 1)
    for (r != nil) {
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
        if err := binary.Read(r, binary.LittleEndian, &item.value); err != nil {
            return nil, err
        }
        
        // Reverse
        // sort.Sort(sort.Reverse(ByteSequence(item.value)))
        items = append(items, item)
    }
    
    return &tlv8Container{
        Items:items,
    }, nil
}

func (t *tlv8Container) GetBuffer(tag uint8) *bytes.Buffer {
    var b bytes.Buffer
    for _, item := range t.Items {
        if item.tag == tag {
            b.Write(item.value)
        }
    }
    
    return &b
}

func (t *tlv8Container) GetString(tag uint8) string {
    return string(t.GetBytes(tag))
}

func (t *tlv8Container) GetBytes(tag uint8) []byte {
    return t.GetBuffer(tag).Bytes()
}

func (t *tlv8Container) GetByte(tag uint8) byte {
    buffer := t.GetBuffer(tag)
    b, _ := buffer.ReadByte()
    return b
}

func (t *tlv8Container) SetString(tag uint8, value string) {
    t.SetBytes(tag, []byte(value))
}

func (t *tlv8Container) SetBytes(tag uint8, value []byte) {
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

func (t *tlv8Container) SetByte(tag uint8, b byte) {
    t.SetBytes(tag, []byte{b})
}

func (t *tlv8Container) BytesBuffer() *bytes.Buffer {
    var b bytes.Buffer
    for _, item := range t.Items {
        // Since we are using just 1 byte for tag and length, the byte order does not matter
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