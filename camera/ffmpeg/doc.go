// Package ffmpeg lets you access the camera via ffmpeg to stream video and to create snapshots.
//
// This package requires the `ffmpeg` command line tool to be installed. Install by running
// - `apt-get install ffmpeg` on linux
// - `brew install ffmpeg` on macOS
//
// HomeKit supports multiple video codecs but h264 is mandatory. So make sure that a h264 decoder for ffmpeg is installed too.
// Audio streaming is currently not supported.
//
// If you are running a RPi with Rasbian, it is recommended to use a v4l2 loopback device instead of access the camera via `/dev/video0` directly.
package ffmpeg
