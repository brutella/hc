package rtp

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/tlv8"
	"testing"
)

func TestStreamController(t *testing.T) {
	c := characteristic.NewSupportedVideoStreamConfiguration()
	c.Value = "AX8BAQACDAEBAQEBAgIBAAMBAAECgAcCAjgEAwEeAQIABQIC0AIDAR4BAoACAgJoAQMBHgEC4AECAg4BAwEeAQJAAQICtAADAR4BAgAFAgLAAwMBHgECAAQCAgADAwEeAQKAAgIC4AEDAR4BAuABAgJoAQMBHgECQAECAvAAAwEP"

	b := c.GetValue()
	if len(b) == 0 {
		t.Fatalf("Zero length bytes")
	}

	var cfg VideoStreamConfiguration
	err := tlv8.Unmarshal(b, &cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMarshalVideoCodecConfiguration(t *testing.T) {
	codec := NewH264VideoCodecConfiguration()
	b, err := tlv8.Marshal(codec)
	if err != nil {
		t.Fatal(err)
	}

	var c VideoCodecConfiguration
	err = tlv8.Unmarshal(b, &c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStreamingStatus(t *testing.T) {
	c := characteristic.NewStreamingStatus()
	c.Value = "AQEA"

	b := c.GetValue()
	if len(b) == 0 {
		t.Fatalf("Zero length bytes")
	}

	var status StreamingStatus
	err := tlv8.Unmarshal(b, &status)
	if err != nil {
		t.Fatal(err)
	}
}
