package pair

import(
    "io"
    "fmt"
)

type PairingHandler interface {
    Handle(Container) (Container, error)
}

func HandleReaderForHandler(r io.Reader, h PairingHandler) (io.Reader, error) {
    cont_in, err := NewTLV8ContainerFromReader(r)
    if err != nil {
        return nil, err
    }
    
    fmt.Println("->     Seq:", cont_in.GetByte(TLVType_SequenceNumber))
    
    cont_out, err := h.Handle(cont_in)
    
    if err != nil {
        fmt.Println("[ERROR]", err)
        return nil, err
    } else {
        if cont_out != nil {
            fmt.Println("<-     Seq:", cont_out.GetByte(TLVType_SequenceNumber))
            fmt.Println("-------------")
            return cont_out.BytesBuffer(), nil
        }
    }
    
    return nil, err
}