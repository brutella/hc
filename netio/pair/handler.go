package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/netio"
	"github.com/brutella/log"

	"io"
)

// HandleReaderForHandler wraps h.Handle() call and logs sequence numbers and errors to the console.
func HandleReaderForHandler(r io.Reader, h netio.ContainerHandler) (r_out io.Reader, err error) {
	cont_in, err := common.NewTLV8ContainerFromReader(r)
	if err != nil {
		return nil, err
	}

	log.Println("[VERB] ->     Seq:", cont_in.GetByte(TagSequence))

	cont_out, err := h.Handle(cont_in)

	if err != nil {
		log.Println("[ERRO]", err)
	} else {
		if cont_out != nil {
			log.Println("[VERB] <-     Seq:", cont_out.GetByte(TagSequence))
			r_out = cont_out.BytesBuffer()
		}
	}
	log.Println("[VERB] --------------------------")

	return r_out, err
}
