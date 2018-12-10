package rtp

const (
	AudioCodecType_PCMU    byte = 0
	AudioCodecType_PCMA         = 1
	AudioCodecType_AAC_ELD      = 2
	AudioCodecType_Opus         = 3
	AudioCodecType_MSBC         = 4
	AudioCodecType_AMR          = 5
	AudioCodecType_ARM_WB       = 6

	AudioCodecBitrateVariable byte = 0
	AudioCodecBitrateConstant byte = 1

	AudioCodecSampleRate8Khz  byte = 0
	AudioCodecSampleRate16Khz byte = 1
	AudioCodecSampleRate24Khz byte = 2
)

type SupportedAudioStreamConfiguration struct {
	Codecs       []AudioCodecConfiguration `tlv8:"1"`
	ComfortNoise bool                      `tlv8:"2"`
}

type AudioCodecConfiguration struct {
	Type       byte                 `tlv8:"1"` // docu says 2 bytes?
	Parameters AudioCodecParameters `tlv8:"2"`
}

type AudioCodecParameters struct {
	Channels   byte `tlv8:"1"` // Default 1
	Bitrate    byte `tlv8:"2"`
	Samplerate byte `tlv8:"3"`

	// "Note: This TLV will only be presented in the Selected Audio"
	// PacketTime byte `tlv8:"4"` // RFC 4566, supported values: 20, 30, 40, 60 ms
}

func DefaultAudioStreamConfiguration() SupportedAudioStreamConfiguration {
	return SupportedAudioStreamConfiguration{
		Codecs: []AudioCodecConfiguration{
			NewOpusAudioCodecConfiguration(),
			NewAacEldAudioCodecConfiguration(),
		},
		ComfortNoise: false,
	}
}

func NewOpusAudioCodecConfiguration() AudioCodecConfiguration {
	return AudioCodecConfiguration{
		Type: AudioCodecType_Opus,
		Parameters: AudioCodecParameters{
			Channels:   1,
			Bitrate:    AudioCodecBitrateVariable,
			Samplerate: AudioCodecSampleRate24Khz,
		},
	}
}

func NewAacEldAudioCodecConfiguration() AudioCodecConfiguration {
	return AudioCodecConfiguration{
		Type: AudioCodecType_AAC_ELD,
		Parameters: AudioCodecParameters{
			Channels:   1,
			Bitrate:    AudioCodecBitrateVariable,
			Samplerate: AudioCodecSampleRate16Khz,
		},
	}
}
