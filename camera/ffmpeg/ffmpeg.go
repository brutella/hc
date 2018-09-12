package ffmpeg

import (
	"fmt"
	"github.com/brutella/hc/camera/rtp"
	"github.com/brutella/hc/log"
	"image"
	"io/ioutil"
	"os"
	"sync"
)

// StreamID is the type of the stream identifier
type StreamID string

// FFMPEG lets you interact with camera stream.
type FFMPEG interface {
	PrepareNewStream(rtp.SetupEndpoints, rtp.SetupEndpointsResponse) StreamID
	Start(StreamID, rtp.VideoParameters, rtp.AudioParameters) error
	Stop(StreamID)
	Suspend(StreamID)
	Resume(StreamID)
	ActiveStreams() int
	Reconfigure(StreamID, rtp.VideoParameters, rtp.AudioParameters) error
	Snapshot(width, height uint) (*image.Image, error)
}

var Stdout = ioutil.Discard
var Stderr = ioutil.Discard

// EnableVerboseLogging enables verbose logging of ffmpeg to stdout.
func EnableVerboseLogging() {
	Stdout = os.Stdout
	Stderr = os.Stderr
}

type ffmpeg struct {
	cfg     Config
	loop    *loopback
	mutex   *sync.Mutex
	streams map[StreamID]*stream
}

// New returns a new ffmpeg handle to start and stop video streams and to make snapshots.
// If cfg specifies a video loopback, ffmpeg configures a loopback to support simultaneous access to the video device.
func New(cfg Config) *ffmpeg {
	var loop *loopback = nil
	if cfg.LoopbackFilename != "" {
		loop = NewLoopback(cfg.InputDevice, cfg.InputFilename, cfg.LoopbackFilename)
	}

	return &ffmpeg{
		cfg:     cfg,
		loop:    loop,
		mutex:   &sync.Mutex{},
		streams: make(map[StreamID]*stream, 0),
	}
}

func (f *ffmpeg) PrepareNewStream(req rtp.SetupEndpoints, resp rtp.SetupEndpointsResponse) StreamID {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	id := StreamID(req.SessionId)
	s := &stream{f.videoInputDevice(), f.videoInputFilename(), f.cfg.H264Decoder, f.cfg.H264Encoder, req, resp, nil}
	f.streams[id] = s
	return id
}

func (f *ffmpeg) ActiveStreams() int {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return len(f.streams)
}

func (f *ffmpeg) Start(id StreamID, video rtp.VideoParameters, audio rtp.AudioParameters) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	s, err := f.getStream(id)
	if err != nil {
		log.Info.Println("start:", err)
		return err
	}

	f.startLoopback()

	return s.start(video, audio)
}

func (f *ffmpeg) Stop(id StreamID) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	s, err := f.getStream(id)
	if err != nil {
		log.Info.Println("stop:", err)
		return
	}

	s.stop()
	delete(f.streams, id)

	if f.loop != nil {
		for _, s := range f.streams {
			if s.isActive() {
				log.Debug.Printf("Active sesssions %v\n", f.streams)
				return
			}
		}

		log.Debug.Println("Stopping loopback")
		// Stop loopback if no stream is active anymore
		f.loop.Stop()
	}
}

func (f *ffmpeg) Suspend(id StreamID) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if s, err := f.getStream(id); err != nil {
		log.Info.Println("suspend:", err)
	} else {
		s.suspend()
	}
}

func (f *ffmpeg) Resume(id StreamID) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if s, err := f.getStream(id); err != nil {
		log.Info.Println("resume:", err)
	} else {
		s.resume()
	}
}

func (f *ffmpeg) Reconfigure(id StreamID, video rtp.VideoParameters, audio rtp.AudioParameters) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	s, err := f.getStream(id)
	if err != nil {
		log.Info.Println("reconfigure:", err)
		return err
	}

	return s.reconfigure(video, audio)
}

func (f *ffmpeg) getStream(id StreamID) (*stream, error) {
	if s, ok := f.streams[id]; ok {
		return s, nil
	}

	return nil, &StreamNotFoundError{id}
}

func (f *ffmpeg) startLoopback() {
	if f.loop != nil {
		if err := f.loop.Start(); err != nil {
			log.Info.Println("starting loopback failed:", err)
		}
	}
}

func (f *ffmpeg) Snapshot(width, height uint) (*image.Image, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.startLoopback()

	shot, err := snapshot(width, height, f.videoInputDevice(), f.videoInputFilename())

	if f.loop != nil {
		for _, s := range f.streams {
			if s.isActive() {
				log.Debug.Printf("Active sesssions %v\n", f.streams)
				return shot, err
			}
		}

		log.Debug.Println("Stopping loopback")
		// Stop loopback if no stream is active anymore
		f.loop.Stop()
	}

	return shot, err
}

func (f *ffmpeg) videoInputDevice() string {
	return f.cfg.InputDevice
}

func (f *ffmpeg) videoInputFilename() string {
	if f.cfg.LoopbackFilename != "" {
		return f.cfg.LoopbackFilename
	}

	return f.cfg.InputFilename
}

type StreamNotFoundError struct {
	id StreamID
}

func (e *StreamNotFoundError) Error() string {
	return fmt.Sprintf("StreamID(%v) not found", e.id)
}
