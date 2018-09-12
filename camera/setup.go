package camera

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/camera/ffmpeg"
	"github.com/brutella/hc/camera/rtp"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/service"
	"github.com/brutella/hc/tlv8"
	"math/rand"
	"net"
)

// SetupFFMPEGStreaming configures a camera to use ffmpeg to stream video.
// The returned handle can be used to interact with the camera (start, stop, take snapshot…).
func SetupFFMPEGStreaming(cam *accessory.Camera, cfg ffmpeg.Config, ips []net.IP) ffmpeg.FFMPEG {
	ff := ffmpeg.New(cfg)

	setupStreamManagement(cam.StreamManagement1, ff, ips, cfg.MultiStream)
	setupStreamManagement(cam.StreamManagement2, ff, ips, cfg.MultiStream)

	return ff
}

func first(ips []net.IP, filter func(net.IP) bool) net.IP {
	for _, ip := range ips {
		if filter(ip) == true {
			return ip
		}
	}

	return nil
}

func setupStreamManagement(m *service.CameraRTPStreamManagement, ff ffmpeg.FFMPEG, ips []net.IP, multiStream bool) {
	status := rtp.StreamingStatus{rtp.StreamingStatusAvailable}
	setTLV8Payload(m.StreamingStatus.Bytes, status)
	setTLV8Payload(m.SupportedRTPConfiguration.Bytes, rtp.NewSupportedRTPConfiguration(rtp.CryptoSuite_AES_CM_128_HMAC_SHA1_80))
	setTLV8Payload(m.SupportedVideoStreamConfiguration.Bytes, rtp.DefaultVideoStreamConfiguration())
	setTLV8Payload(m.SupportedAudioStreamConfiguration.Bytes, rtp.DefaultAudioStreamConfiguration())

	m.SelectedRTPStreamConfiguration.OnValueRemoteUpdate(func(buf []byte) {
		var cfg rtp.SelectedRtpStreamConfiguration
		err := tlv8.Unmarshal(buf, &cfg)
		if err != nil {
			log.Debug.Fatalf("SelectedRTPStreamConfiguration: Could not unmarshal tlv8 data: %s\n", err)
		}

		log.Debug.Printf("%+v\n", cfg)

		id := ffmpeg.StreamID(cfg.Command.Identifier)
		switch cfg.Command.Type {
		case rtp.SessionControlCommandTypeEnd:
			ff.Stop(id)

			if ff.ActiveStreams() == 0 {
				// Update stream status when no streams are currently active
				setTLV8Payload(m.StreamingStatus.Bytes, rtp.StreamingStatus{rtp.StreamingStatusAvailable})
			}

		case rtp.SessionControlCommandTypeStart:
			ff.Start(id, cfg.Video, cfg.Audio)

			if multiStream == false {
				// If only one video stream is suppported, set the status to busy.
				// This way HomeKit knows that nobody is allowed to connect anymore.
				// If multiple streams are supported, the status is always availabe.
				setTLV8Payload(m.StreamingStatus.Bytes, rtp.StreamingStatus{rtp.StreamingStatusBusy})
			}
		case rtp.SessionControlCommandTypeSuspend:
			ff.Suspend(id)
		case rtp.SessionControlCommandTypeResume:
			ff.Resume(id)
		case rtp.SessionControlCommandTypeReconfigure:
			ff.Reconfigure(id, cfg.Video, cfg.Audio)
		default:
			log.Debug.Printf("Unknown command type %d", cfg.Command.Type)
		}
	})

	m.SetupEndpoints.OnValueRemoteUpdate(func(buf []byte) {
		var req rtp.SetupEndpoints
		err := tlv8.Unmarshal(buf, &req)
		if err != nil {
			log.Debug.Fatalf("SetupEndpoints: Could not unmarshal tlv8 data: %s\n", err)
		}

		log.Debug.Printf("%+v\n", req)

		// find the ip adress with the same version as the controller
		ip := first(ips, func(ip net.IP) bool {
			switch req.ControllerAddr.IPVersion {
			case rtp.IPAddrVersionv4:
				return ip.To4() != nil
			case rtp.IPAddrVersionv6:
				return ip.To4() == nil
			default:
				return false
			}
		})

		if ip == nil {
			log.Info.Println("No IP address of version", req.ControllerAddr.IPVersion)
			return
		}

		var version = rtp.IPAddrVersionv4
		if ip.To4() == nil {
			version = rtp.IPAddrVersionv6
		}

		// ssrc is different for every stream
		ssrcVideo := rand.Int31n(1000)
		ssrcAudio := rand.Int31n(1000)

		resp := rtp.SetupEndpointsResponse{
			SessionId: req.SessionId,
			Status:    rtp.SessionStatusSuccess,
			AccessoryAddr: rtp.Addr{
				IPVersion:    version,
				IPAddr:       ip.String(),
				VideoRtpPort: req.ControllerAddr.VideoRtpPort,
				AudioRtpPort: req.ControllerAddr.AudioRtpPort,
			},
			Video:     req.Video,
			Audio:     req.Audio,
			SsrcVideo: ssrcVideo,
			SsrcAudio: ssrcAudio,
		}

		ff.PrepareNewStream(req, resp)

		log.Debug.Printf("%+v\n", resp)

		// After a write, the characteristic should contain a response
		setTLV8Payload(m.SetupEndpoints.Bytes, resp)
	})
}

func setTLV8Payload(c *characteristic.Bytes, v interface{}) {
	if tlv8, err := tlv8.Marshal(v); err == nil {
		c.SetValue(tlv8)
	} else {
		log.Debug.Fatal(err)
	}
}
