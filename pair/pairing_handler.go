package pair

import(
    "io"
    "fmt"
)

type PairingHandler interface {
    Handle(Container) (Container, error)
}

func HandleReaderForHandler(r io.Reader, h PairingHandler) (r_out io.Reader, err error) {
    cont_in, err := NewTLV8ContainerFromReader(r)
    if err != nil {
        return nil, err
    }
    
    fmt.Println("->     Seq:", cont_in.GetByte(TLVType_SequenceNumber))
    
    cont_out, err := h.Handle(cont_in)
    
    if err != nil {
        fmt.Println("[ERROR]", err)
    } else {
        if cont_out != nil {
            fmt.Println("<-     Seq:", cont_out.GetByte(TLVType_SequenceNumber))
            r_out = cont_out.BytesBuffer()
        }
    }
    fmt.Println("--------------------------")
    
    return r_out, err
}