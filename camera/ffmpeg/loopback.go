package ffmpeg

import (
	"os/exec"
	"syscall"
	"time"
)

// loopback copies data from the inpute filename to the loopback filename.
// This is needed to provide simultaneous access to a v4l2 device.
// On a Raspberry Pi you can create a video loopback device at /dev/video1 via [v4l2loopback](https://github.com/umlaeute/v4l2loopback).
// The data from /dev/video0 is then copied to /dev/video1 via `ffmpeg -f v4l2 -i /dev/video0 -codec:v copy -f v4l2 /dev/video1`.
// The loopback device can then be access simultanously from multiple ffmpeg processes.
type loopback struct {
	inputDevice      string
	inputFilename    string
	h264Decoder      string
	loopbackFilename string

	active *exec.Cmd
}

// NewLoopback returns a new video loopback.
func NewLoopback(inputDevice, inputFilename, loopbackFilename string) *loopback {
	return &loopback{
		inputDevice:      inputDevice,
		inputFilename:    inputFilename,
		loopbackFilename: loopbackFilename,
	}
}

// Start starts the loopback.
func (l *loopback) Start() error {
	if l.active == nil {
		cmd := l.cmd()
		if err := cmd.Start(); err != nil {
			return err
		}
		// FIXME: Starting ffmpeg takes some time
		time.Sleep(2 * time.Second)
		l.active = cmd
	}

	return nil
}

// Stop stops the loopback.
func (l *loopback) Stop() {
	if l.active != nil {
		l.active.Process.Signal(syscall.SIGINT)
		l.active = nil
	}
}

// cmd returns a new command to stream video from the input file to the loopback file.
func (l *loopback) cmd() *exec.Cmd {
	cmd := exec.Command("ffmpeg", "-f", l.inputDevice, "-i", l.inputFilename, "-codec:v", "copy", "-f", l.inputDevice, l.loopbackFilename)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr

	return cmd
}
