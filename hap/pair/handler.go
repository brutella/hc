package pair

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/util"

	"io"
)

// HandleReaderForHandler wraps h.Handle() call and logs sequence numbers and errors to the console.
func HandleReaderForHandler(r io.Reader, h hap.ContainerHandler) (rOut io.Reader, err error) {
	in, err := util.NewTLV8ContainerFromReader(r)
	if err != nil {
		return nil, err
	}

	log.Debug.Println("->     Seq:", in.GetByte(TagSequence))

	out, err := h.Handle(in)

	if err != nil {
		log.Info.Println(err)
	} else {
		if out != nil {
			log.Debug.Println("<-     Seq:", out.GetByte(TagSequence))
			rOut = out.BytesBuffer()
		}
	}
	log.Debug.Println("--------------------------")

	return rOut, err
}
