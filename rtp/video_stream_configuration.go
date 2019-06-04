package rtp

const (
	VideoCodecType_H264 byte = 0

	VideoCodecProfileConstrainedBaseline byte = 0
	VideoCodecProfileMain                     = 1
	VideoCodecProfileHigh                     = 2

	VideoCodecLevel3_1 byte = 0
	VideoCodecLevel3_2 byte = 1
	VideoCodecLevel4   byte = 2

	VideoCodecPacketizationModeNonInterleaved byte = 0

	VideoCodecCvoNotSuppported byte = 0
	VideoCodecCvoSuppported         = 1
)

type VideoStreamConfiguration struct {
	Codecs []VideoCodecConfiguration `tlv8:"1"`
}

type VideoCodecConfiguration struct {
	Type       byte                   `tlv8:"1"`
	Parameters VideoCodecParameters   `tlv8:"2"`
	Attributes []VideoCodecAttributes `tlv8:"3"`
}

type VideoCodecParameters struct {
	Profiles       []VideoCodecProfile       `tlv8:"-"`
	Levels         []VideoCodecLevel         `tlv8:"-"`
	Packetizations []VideoCodecPacketization `tlv8:"-"`
	// CvoEnabled     bool                      `-tlv8:"4,optional"` // ?
	// CvoId       byte                      `-tlv8:"5,optional"` // ? value from [1:14]
}

type VideoCodecProfile struct {
	Id byte `tlv8:"1"`
}

type VideoCodecLevel struct {
	Level byte `tlv8:"2"`
}

type VideoCodecPacketization struct {
	Mode byte `tlv8:"3"`
}

type VideoCodecAttributes struct {
	Width     uint16 `tlv8:"1"`
	Height    uint16 `tlv8:"2"`
	Framerate byte   `tlv8:"3"`
}

func DefaultVideoStreamConfiguration() VideoStreamConfiguration {
	return VideoStreamConfiguration{
		Codecs: []VideoCodecConfiguration{
			NewH264VideoCodecConfiguration(),
		},
	}
}

func NewH264VideoCodecConfiguration() VideoCodecConfiguration {
	return VideoCodecConfiguration{
		Type: VideoCodecType_H264,
		Parameters: VideoCodecParameters{
			Profiles: []VideoCodecProfile{
				VideoCodecProfile{VideoCodecProfileConstrainedBaseline},
				VideoCodecProfile{VideoCodecProfileMain},
				VideoCodecProfile{VideoCodecProfileHigh},
			},
			Levels: []VideoCodecLevel{
				VideoCodecLevel{VideoCodecLevel3_1},
				VideoCodecLevel{VideoCodecLevel3_2},
				VideoCodecLevel{VideoCodecLevel4},
			},
			Packetizations: []VideoCodecPacketization{
				VideoCodecPacketization{VideoCodecPacketizationModeNonInterleaved},
			},
		},
		Attributes: []VideoCodecAttributes{
			VideoCodecAttributes{1920, 1080, 30}, // 1080p
			VideoCodecAttributes{1280, 720, 30},  // 720p
			VideoCodecAttributes{640, 360, 30},
			VideoCodecAttributes{480, 270, 30},
			VideoCodecAttributes{320, 180, 30},
			VideoCodecAttributes{1280, 960, 30},
			VideoCodecAttributes{1024, 768, 30}, // XVGA
			VideoCodecAttributes{640, 480, 30},  // VGA
			VideoCodecAttributes{480, 360, 30},  // HVGA
			VideoCodecAttributes{320, 240, 15},  // QVGA (Apple Watch)
		},
	}
}
