package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/log"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
)

type SnapshotFunc func(width, height uint) (*image.Image, error)

// Resource handles the /resource endpoint
type Resource struct {
	http.Handler
	snapshot SnapshotFunc
	context  hap.Context
}

// NewResource returns a new handler for resource requests
func NewResource(context hap.Context, snapshot SnapshotFunc) *Resource {
	r := Resource{
		context:  context,
		snapshot: snapshot,
	}

	return &r
}

type ImageRequest struct {
	Type   string `json:"resource-type"`
	Width  uint   `json:"image-width"`
	Height uint   `json:"image-height"`
}

func (handler *Resource) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case hap.MethodPOST:
		log.Debug.Printf("%v POST /resource\n", request.RemoteAddr)
		if err := handler.postResource(response, request); err != nil {
			log.Info.Println(err)
			response.WriteHeader(http.StatusInternalServerError)
		}
	default:
		log.Debug.Println("Cannot handle HTTP method", request.Method)
		response.WriteHeader(http.StatusNoContent)
	}
}

func (r *Resource) postResource(resp http.ResponseWriter, req *http.Request) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var imgRq ImageRequest
	err = json.Unmarshal(body, &imgRq)
	if err != nil {
		return err
	}

	if imgRq.Type == "image" {
		b, err := r.getJPEGSnapshot(imgRq)
		if err != nil {
			return err
		}

		resp.Header().Set("Content-Type", "image/jpeg")
		wr := hap.NewChunkedWriter(resp, 2048)
		wr.Write(b)
		return nil
	}

	resp.WriteHeader(http.StatusNoContent)

	return nil
}

func (r *Resource) getJPEGSnapshot(req ImageRequest) ([]byte, error) {
	img, err := r.snapshot(req.Width, req.Height)
	if err != nil {
		return nil, fmt.Errorf("r.snapshot() %v", err)
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, *img, nil); err != nil {
		return nil, fmt.Errorf("jpeg.Encode() %v", err)
	}

	return buf.Bytes(), nil
}
