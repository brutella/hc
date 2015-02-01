package pair

import (
	"github.com/brutella/hap/common"
	"github.com/brutella/hap/netio"
	"github.com/brutella/log"

	"io"
)

func HandleReaderForHandler(r io.Reader, h netio.ContainerHandler) (r_out io.Reader, err error) {
	cont_in, err := common.NewTLV8ContainerFromReader(r)
	if err != nil {
		return nil, err
	}

	log.Println("[VERB] ->     Seq:", cont_in.GetByte(TLVType_SequenceNumber))

	cont_out, err := h.Handle(cont_in)

	if err != nil {
		log.Println("[ERRO]", err)
	} else {
		if cont_out != nil {
			log.Println("[VERB] <-     Seq:", cont_out.GetByte(TLVType_SequenceNumber))
			r_out = cont_out.BytesBuffer()
		}
	}
	log.Println("[VERB] --------------------------")

	return r_out, err
}
