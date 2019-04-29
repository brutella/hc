package rtp

type StreamingStatus struct {
	Status byte `tlv8:"1"`
}

const (
	StreamingStatusAvailable   byte = 0
	StreamingStatusBusy        byte = 1
	StreamingStatusUnavailable byte = 2
)
