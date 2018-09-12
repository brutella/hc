package ffmpeg

// Config contains ffmpeg parameters
type Config struct {
	InputDevice      string
	InputFilename    string
	LoopbackFilename string
	H264Decoder      string
	H264Encoder      string
	MultiStream      bool
}
