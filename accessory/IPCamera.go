package accessory

import (
	"github.com/brutella/hc/service"
	"github.com/brutella/hc/util"
)

//IPCamera struct
type IPCamera struct {
	*Accessory
	CameraControl             *service.CameraControl
	CameraRTPStreamManagement *service.CameraRTPStreamManagement
}

// NewIPCamera func
func NewIPCamera(info Info) *IPCamera {
	acc := IPCamera{}
	acc.Accessory = New(info, TypeIPCamera)

	acc.CameraControl = service.NewCameraControl()
	acc.AddService(acc.CameraControl.Service)

	// Configure and add Camera RTP Stream Management service
	acc.CameraRTPStreamManagement = service.NewCameraRTPStreamManagement()

	// Setup Supported Video Stream Configuration
	supportedVideoConfigurationTlv := util.NewTLV8Container()
	supportedVideoConfigurationTlv.SetByte(1, 0) // H.264
	tlvCodecParams := util.NewTLV8Container()
	tlvCodecParams.SetByte(1, 1) // ProfileID: Main Profile
	tlvCodecParams.SetByte(2, 2) // Level: 4
	tlvCodecParams.SetByte(3, 0) // Packetization mode: Non-interleaved mode
	tlvCodecParams.SetByte(4, 0) // CVO Enabled: CVO not supported
	supportedVideoConfigurationTlv.SetBytes(2, tlvCodecParams.BytesBuffer().Bytes())

	tlvVideoAttrs := util.NewTLV8Container()
	tlvVideoAttrs.SetBytes(1, []byte{128, 7}) // Image width: 1920
	tlvVideoAttrs.SetBytes(2, []byte{56, 4})  // Image height: 1080
	tlvVideoAttrs.SetByte(3, 30)              // Frame rate: 30
	supportedVideoConfigurationTlv.SetBytes(3, tlvVideoAttrs.BytesBuffer().Bytes())

	acc.CameraRTPStreamManagement.SupportedVideoStreamConfiguration.SetValue(supportedVideoConfigurationTlv.BytesBuffer().Bytes())

	// Setup Supported Audio Stream Configuration
	supportedAudioConfigurationTlv := util.NewTLV8Container()
	tlvAudioCodecConfig := util.NewTLV8Container()
	tlvAudioCodecConfig.SetByte(1, 3) // Codec type: Opus
	tlvAudioCodecParams := util.NewTLV8Container()
	tlvAudioCodecParams.SetByte(1, 1) // Channels: 1
	tlvAudioCodecParams.SetByte(2, 0) // Bit-rate: VBR
	tlvAudioCodecParams.SetByte(3, 2) // Sample rate: 24KHz
	tlvAudioCodecConfig.SetBytes(2, tlvAudioCodecParams.BytesBuffer().Bytes())
	tlvAudioCodecConfig.SetByte(2, 0) // Comfort Noise support: false
	supportedAudioConfigurationTlv.SetBytes(1, tlvAudioCodecConfig.BytesBuffer().Bytes())

	// Setup Supported RTP Configuration
	supportedRTPConfigurationTlv := util.NewTLV8Container()
	supportedRTPConfigurationTlv.SetByte(2, 2) // SRTP Crypto Suite: Disabled
	acc.CameraRTPStreamManagement.SupportedRTPConfiguration.SetValue(supportedRTPConfigurationTlv.BytesBuffer().Bytes())

	// Setup Streaming Status
	streamingStatusTlv := util.NewTLV8Container()
	streamingStatusTlv.SetByte(1, 0) // Status: Available
	acc.CameraRTPStreamManagement.StreamingStatus.SetValue(streamingStatusTlv.BytesBuffer().Bytes())

	// Add Camera RTP Stream Management service to accessory
	acc.AddService(acc.CameraRTPStreamManagement.Service)

	return &acc
}
