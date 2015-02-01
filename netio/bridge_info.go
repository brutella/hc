package netio

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/brutella/hap/common"
	"strings"
)

type BridgeInfo struct {
	SerialNumber string
	Password     string
	Name         string
	Id           string
	Manufacturer string
}

func NewBridgeInfo(name, password, manufacturer string, storage common.Storage) BridgeInfo {
	serial := common.GetSerialNumberForAccessoryName(name, storage)
	return BridgeInfo{
		SerialNumber: serial,
		Password:     password,
		Name:         name,
		Id:           MAC48Address(serial),
		Manufacturer: manufacturer,
	}
}

// Returns a MAC-48 address
func MAC48Address(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	result := h.Sum(nil)

	c := make([]string, 0)
	c = append(c, toHex(result[0]))
	c = append(c, toHex(result[1]))
	c = append(c, toHex(result[2]))
	c = append(c, toHex(result[3]))
	c = append(c, toHex(result[4]))
	c = append(c, toHex(result[5]))

	return strings.Join(c, ":")
}

func toHex(b byte) string {
	return hex.EncodeToString([]byte{b})
}
