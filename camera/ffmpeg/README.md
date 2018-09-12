# Debugging
ffmpeg can be used to debug SRTP streaming by using ffmpeg and ffplay.
Install ffmpeg and ffplay via `brew install ffmpeg --with-ffplay`.

## RTSP Streaming

1. Create an RTSP stream
1.1. Raspberry Pi
    ffmpeg -re -f video4linux2 -i /dev/video0 -map 0:0 -vcodec h264_omx -pix_fmt yuv420p -r 20 -f rawvideo -tune zerolatency -b:v 1500k -bufsize 1500k -payload_type 99 -ssrc 16132552 -f rtp -srtp_out_suite AES_CM_128_HMAC_SHA1_80 -srtp_out_params omz31e5SiZSneUySvSsIaFfu+NW2uWUl9+FHs3HD "srtp://192.168.0.14:58536?rtcpport=58536&localrtcpport=58536&pkt_size=1378"
1.2. Mac
    ffmpeg -re -f avfoundation -i "1" -map 0:0 -vcodec libx264 -pix_fmt yuv420p -r 20 -f rawvideo -tune zerolatency -b:v 1500k -bufsize 1500k -payload_type 99 -ssrc 16132552 -f rtp -srtp_out_suite AES_CM_128_HMAC_SHA1_80 -srtp_out_params omz31e5SiZSneUySvSsIaFfu+NW2uWUl9+FHs3HD "srtp://192.168.0.14:58536?rtcpport=58536&localrtcpport=58536&pkt_size=1378"

2. Create sdp file on receiver
    v=0
    o=- 0 0 IN IP4 127.0.0.1
    s=No Name
    c=IN IP4 192.168.0.14
    t=0 0
    a=tool:libavformat 58.17.101
    m=video 58536 RTP/AVP 99
    b=AS:300
    a=rtpmap:99 H264/90000
    a=fmtp:99 packetization-mode=1
    a=crypto:1 AES_CM_128_HMAC_SHA1_80 inline:omz31e5SiZSneUySvSsIaFfu+NW2uWUl9+FHs3HD

3. Receive stream
    ffplay -i <path-to-sdp-file> -protocol_whitelist file,udp,rtp

## Bitrate
FFMPEG cannot set the bitrate on Rasbperry Pi, we have to do it manually with `v4l2-ctl --set-ctrl video_bitrate=300000`.

## Streaming issues
ffmpeg only sends one H264 keyframe at the beginning of a RTP stream on the RPi because of https://video.stackexchange.com/a/21245
If we open the stream after it has been started using ffplay, we missed the keyframe and never get one – results in error: decode_slice_header error.
Therefore we have to start ffplay before starting streaming.