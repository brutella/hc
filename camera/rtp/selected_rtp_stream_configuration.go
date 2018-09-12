package rtp

type SelectedRtpStreamConfiguration struct {
	Command SessionControlCommand `tlv8:"1"`
	Video   VideoParameters       `tlv8:"2"`
	Audio   AudioParameters       `tlv8:"3"`
}

const (
	SessionControlCommandTypeEnd         byte = 0
	SessionControlCommandTypeStart            = 1
	SessionControlCommandTypeSuspend          = 2
	SessionControlCommandTypeResume           = 3
	SessionControlCommandTypeReconfigure      = 4
)

type SessionControlCommand struct {
	Identifier []byte `tlv8:"1"` // docu says 16 bytes but overall size of SessionControlCommand is also 16 bytes?
	Type       byte   `tlv8:"2"`
}

type VideoParameters struct {
	CodecType   byte                 `tlv8:"1"`
	CodecParams VideoCodecParameters `tlv8:"2"`
	Attributes  VideoCodecAttributes `tlv8:"3"`
	RTP         RTPParams            `tlv8:"4"`
}

type AudioParameters struct {
	CodecType    byte                 `tlv8:"1"`
	CodecParams  AudioCodecParameters `tlv8:"2"`
	RTP          RTPParams            `tlv8:"3"`
	ComfortNoise bool                 `tlv8:"4"`
}

type RTPParams struct {
	PayloadType             uint8   `tlv8:"1"`
	Ssrc                    int32   `tlv8:"2"`
	Bitrate                 uint16  `tlv8:"3"`
	Interval                float32 `tlv8:"4"` // MinimumRTCP interval
	ComfortNoisePayloadType uint8   `tlv8:"5"` // only for audio
	MTU                     uint16  `tlv8:"6"` // only for video
}
