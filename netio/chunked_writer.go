package netio

import (
	"io"
)

type chunkedWriter struct {
	wr    io.Writer
	chunk int
}

// NewChunkedWriter returns a writer which writes bytes in chunkes of specified size.
func NewChunkedWriter(wr io.Writer, chunk int) io.Writer {
	return &chunkedWriter{wr, chunk}
}

func (w *chunkedWriter) Write(p []byte) (int, error) {
	var max int = len(p)

	var nn int = 0
	var end int
	for nn < max {
		end = nn + w.chunk
		if end > max {
			end = max
		}
		n, err := w.wr.Write(p[nn:end])
		if err != nil {
			return nn, err
		}
		nn += n
	}

	return nn, nil
}
