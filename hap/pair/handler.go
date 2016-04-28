package pair

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/util"
	"github.com/brutella/log"

	"io"
)

// HandleReaderForHandler wraps h.Handle() call and logs sequence numbers and errors to the console.
func HandleReaderForHandler(r io.Reader, h hap.ContainerHandler) (rOut io.Reader, err error) {
	in, err := util.NewTLV8ContainerFromReader(r)
	if err != nil {
		return nil, err
	}

	log.Println("[VERB] ->     Seq:", in.GetByte(TagSequence))

	out, err := h.Handle(in)

	if err != nil {
		log.Println("[ERRO]", err)
	} else {
		if out != nil {
			log.Println("[VERB] <-     Seq:", out.GetByte(TagSequence))
			rOut = out.BytesBuffer()
		}
	}
	log.Println("[VERB] --------------------------")

	return rOut, err
}
