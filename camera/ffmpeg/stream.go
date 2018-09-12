package ffmpeg

import (
	"fmt"
	"github.com/brutella/hc/camera/rtp"
	"github.com/brutella/hc/log"
	"os/exec"
	"strings"
	"syscall"
)

type stream struct {
	inputDevice   string
	inputFilename string
	h264Decoder   string
	h264Encoder   string
	req           rtp.SetupEndpoints
	resp          rtp.SetupEndpointsResponse

	cmd *exec.Cmd
}

func (s *stream) isActive() bool {
	return s.cmd != nil
}

func (s *stream) stop() {
	log.Debug.Println("stop stream")

	if s.cmd != nil {
		s.cmd.Process.Signal(syscall.SIGINT)
		s.cmd = nil
	}
}

func (s *stream) start(video rtp.VideoParameters, audio rtp.AudioParameters) error {
	log.Debug.Println("start stream")

	// -vsync 2: Fixes "Frame rate very high for a muxer not efficiently supporting it."
	// -framerate before -i specifies the framerate for the input, after -i sets it for the output https://stackoverflow.com/questions/38498599/webcam-with-ffmpeg-on-mac-selected-framerate-29-970030-is-not-supported-by-th#38549528

	// ffmpeg -i input.jpg -vf scale=w=320:h=240:force_original_aspect_ratio=decrease output_320.png
	ffmpegVideo := fmt.Sprintf("-f %s", s.inputDevice) +
		fmt.Sprintf(" -framerate %d", s.framerate(video.Attributes)) +
		fmt.Sprintf("%s", s.videoDecoderOption(video)) +
		fmt.Sprintf(" -i %s", s.inputFilename) +
		" -an" +
		fmt.Sprintf(" -codec:v %s", s.videoEncoder(video)) +
		" -pix_fmt yuv420p -vsync 2" +
		fmt.Sprintf(" -video_size %dx%d", video.Attributes.Width, video.Attributes.Height) +
		fmt.Sprintf(" -framerate %d", video.Attributes.Framerate) +
		// 2018-08-18 (mah) Disable profile arguments because it cannot be parsed
		// [h264_omx @ 0x93a410] [Eval @ 0xbeaad160] Undefined constant or missing '(' in 'high'
		// fmt.Sprintf(" -profile:v %s", videoProfile(video.CodecParams)) +
		fmt.Sprintf(" -level:v %s", videoLevel(video.CodecParams)) +
		" -f rawvideo -tune zerolatency" +
		fmt.Sprintf(" -b:v %dk -bufsize %dk", video.RTP.Bitrate, video.RTP.Bitrate) +
		fmt.Sprintf(" -payload_type %d", video.RTP.PayloadType) +
		fmt.Sprintf(" -ssrc %d", s.resp.SsrcVideo) +
		" -f rtp -srtp_out_suite AES_CM_128_HMAC_SHA1_80" +
		fmt.Sprintf(" -srtp_out_params %s", s.req.Video.SrtpKey()) +
		fmt.Sprintf(" srtp://%s:%d?rtcpport=%d&localrtcpport=%d&pkt_size=%s&timeout=60", s.req.ControllerAddr.IPAddr, s.req.ControllerAddr.VideoRtpPort, s.req.ControllerAddr.VideoRtpPort, s.req.ControllerAddr.VideoRtpPort, videoMTU(s.req))

		// FIXME (mah) Audio doesn't work yet
		// ffmpegAudio := "-vn" +
		//     fmt.Sprintf(" %s", audioCodecOption(audio)) +
		//     // compression-level 0-10 (fastest-slowest)
		//     fmt.Sprintf(" -b:a %dk -bufsize 48k", audio.RTP.Bitrate) +
		//     fmt.Sprintf(" -ar %s", audioSamplingRate(audio)) +
		//     fmt.Sprintf(" -payload_type %d", audio.RTP.PayloadType) +
		// fmt.Sprintf(" -ssrc %d", s.resp.SsrcAudio) +
		//     " -f rtp -srtp_out_suite AES_CM_128_HMAC_SHA1_80" +
		//     fmt.Sprintf(" -srtp_out_params %s", s.req.Audio.SrtpKey()) +
		//     fmt.Sprintf(" srtp://%s:%d?rtcpport=%d&localrtcpport=%d&timeout=60", s.req.ControllerAddr.IPAddr, s.req.ControllerAddr.AudioRtpPort, s.req.ControllerAddr.AudioRtpPort, s.req.ControllerAddr.AudioRtpPort)

	args := strings.Split(ffmpegVideo, " ")
	cmd := exec.Command("ffmpeg", args[:]...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr

	log.Debug.Println(cmd)

	err := cmd.Start()
	if err == nil {
		s.cmd = cmd
	}

	return err
}

// TODO (mah) test
func (s *stream) suspend() {
	log.Debug.Println("suspend stream")
	s.cmd.Process.Signal(syscall.SIGSTOP)
}

// TODO (mah) test
func (s *stream) resume() {
	log.Debug.Println("resume stream")
	s.cmd.Process.Signal(syscall.SIGCONT)
}

// TODO (mah) implement
func (s *stream) reconfigure(video rtp.VideoParameters, audio rtp.AudioParameters) error {
	if s.cmd != nil {
		log.Debug.Println("reconfigure() is not implemented")
	}

	return nil
}

func (s *stream) videoEncoder(param rtp.VideoParameters) string {
	switch param.CodecType {
	case rtp.VideoCodecType_H264:
		return s.h264Encoder
	}

	return "?"
}

func (s *stream) videoDecoderOption(param rtp.VideoParameters) string {
	switch param.CodecType {
	case rtp.VideoCodecType_H264:
		if s.h264Decoder != "" {
			return fmt.Sprintf(" -codec:v %s", s.h264Decoder)
		}
	}

	return ""
}

// https://superuser.com/a/564007
func videoProfile(param rtp.VideoCodecParameters) string {
	for _, p := range param.Profiles {
		switch p.Id {
		case rtp.VideoCodecProfileConstrainedBaseline:
			return "baseline"
		case rtp.VideoCodecProfileMain:
			return "main"
		case rtp.VideoCodecProfileHigh:
			return "high"
		default:
			break
		}
	}

	return ""
}

func (s *stream) framerate(attr rtp.VideoCodecAttributes) byte {
	if s.inputDevice == "avfoundation" {
		// avfoundation only supports 30 fps on a MacBook Pro (Retina, 15-inch, Late 2013) running macOS 10.12 Sierra
		// TODO (mah) test this with other Macs
		return 30
	}

	return attr.Framerate
}

// https://superuser.com/a/564007
func videoLevel(param rtp.VideoCodecParameters) string {
	for _, l := range param.Levels {
		switch l.Level {
		case rtp.VideoCodecLevel3_1:
			return "3.1"
		case rtp.VideoCodecLevel3_2:
			return "3.2"
		case rtp.VideoCodecLevel4:
			return "4.0"
		default:
			break
		}
	}

	return ""
}

func videoMTU(setup rtp.SetupEndpoints) string {
	switch setup.ControllerAddr.IPVersion {
	case rtp.IPAddrVersionv4:
		return "1378"
	case rtp.IPAddrVersionv6:
		return "1228"
	}

	return "1378"
}

// https://trac.ffmpeg.org/wiki/audio%20types
func audioCodecOption(param rtp.AudioParameters) string {
	switch param.CodecType {
	case rtp.AudioCodecType_PCMU:
		log.Debug.Println("audioCodec(PCMU) not supported")
	case rtp.AudioCodecType_PCMA:
		log.Debug.Println("audioCodec(PCMA) not supported")
	case rtp.AudioCodecType_AAC_ELD:
		return "-acodec aac"
		// return "-acodec libfdk_aac -aprofile aac_eld" // requires ffmpeg built with --enable-libfdk-aac
	case rtp.AudioCodecType_Opus:
		// requires ffmpeg built with --enable-libopus
		// - macOS: brew reinstall ffmpeg --with-opus
		return fmt.Sprintf("-acodec libopus")
	case rtp.AudioCodecType_MSBC:
		log.Debug.Println("audioCodec(MSBC) not supported")
	case rtp.AudioCodecType_AMR:
		log.Debug.Println("audioCodec(AMR) not supported")
	case rtp.AudioCodecType_ARM_WB:
		log.Debug.Println("audioCodec(ARM_WB) not supported")
	}

	return ""
}

func audioVariableBitrate(param rtp.AudioParameters) string {
	switch param.CodecParams.Bitrate {
	case rtp.AudioCodecBitrateVariable:
		return "on"
	case rtp.AudioCodecBitrateConstant:
		return "off"
	default:
		log.Info.Println("variableBitrate() undefined bitrate", param.CodecParams.Bitrate)
		break
	}

	return "?"
}

func audioSamplingRate(param rtp.AudioParameters) string {
	switch param.CodecParams.Samplerate {
	case rtp.AudioCodecSampleRate8Khz:
		return "8k"
	case rtp.AudioCodecSampleRate16Khz:
		return "16k"
	case rtp.AudioCodecSampleRate24Khz:
		return "24k"
	default:
		log.Info.Println("audioSamplingRate() undefined samplrate", param.CodecParams.Samplerate)
		break
	}

	return ""
}
