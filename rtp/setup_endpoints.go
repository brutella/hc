package rtp

import (
	"encoding/base64"
	"fmt"
)

type SetupEndpoints struct {
	SessionId      []byte      `tlv8:"1"`
	ControllerAddr Addr        `tlv8:"3"`
	Video          CryptoSuite `tlv8:"4"`
	Audio          CryptoSuite `tlv8:"5"`
}

func (s *SetupEndpoints) String() string {
	return fmt.Sprintf("SessionId: %v\nController\n\tAddr:%v\n\tVideo Port:%v\n\tAudio Port:%v", s.SessionId, s.ControllerAddr.IPAddr, s.ControllerAddr.VideoRtpPort, s.ControllerAddr.AudioRtpPort)
}

type SetupEndpointsResponse struct {
	SessionId     []byte      `tlv8:"1"`
	Status        byte        `tlv8:"2"`
	AccessoryAddr Addr        `tlv8:"3"` // AccessoryAddr.IPVersion must be the same as in SetupEndpoints.ControllerAddr.IPVersion
	Video         CryptoSuite `tlv8:"4"`
	Audio         CryptoSuite `tlv8:"5"`
	SsrcVideo     int32       `tlv8:"6"`
	SsrcAudio     int32       `tlv8:"7"`
}

const (
	SessionStatusSuccess byte = 0
	SessionStatusBusy    byte = 1
	SessionStatusError   byte = 2
)

type Addr struct {
	IPVersion    byte   `tlv8:"1"`
	IPAddr       string `tlv8:"2"`
	VideoRtpPort uint16 `tlv8:"3"`
	AudioRtpPort uint16 `tlv8:"4"`
}

type CryptoSuite struct {
	Types      []CryptoSuiteType `tlv8:"-"`
	MasterKey  []byte            `tlv8:"2"` // 16 (AES_CM_128) or 32 (AES_256_CM)
	MasterSalt []byte            `tlv8:"3"` // 14 byte
}

func (c *CryptoSuite) SrtpKey() string {
	key := append(c.MasterKey, c.MasterSalt[:]...)
	return base64.StdEncoding.EncodeToString(key)
}

type CryptoSuiteType struct {
	Type byte `tlv8:"1"`
}

const (
	IPAddrVersionv4 byte = 0
	IPAddrVersionv6 byte = 1
)
