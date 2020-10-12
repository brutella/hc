package http

import (
	"github.com/brutella/hc/hap"

	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func JSONEncode(v interface{}) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	err := enc.Encode(v)

	return buf, err
}

func JSONDecode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func WriteJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	buf, err := JSONEncode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	wr := hap.NewChunkedWriter(w, 2048)
	wr.Write(buf.Bytes())

	return nil
}

func ReadJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := JSONDecode(r.Body, v); err != nil {
		return err
	}

	return nil
}
