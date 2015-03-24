package netio

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/brutella/hc/common"
	"strings"
)

// BridgeInfo contains all informations to publish a HomeKit bridge
type BridgeInfo struct {
	// The serial number which appears in the bridge's accessory information service
	SerialNumber string
	// The name which appears in the bridge's accessory information service
	Name string
	// The manufacturer name which appears in the bridge's accessory information service
	Manufacturer string
	// The password the user has to enter when adding the accessory to HomeKit
	Password string
	// The id which appears inthe mDNS TXT entry
	ID string
}

// NewBridgeInfo returns a BridgeInfo object.
//
// The BridgeInfo.SerialNumber is loaded from the storage, or created if not found.
// The BridgeInfo.ID is based on the serial number bytes.
func NewBridgeInfo(name, password, manufacturer string, storage common.Storage) BridgeInfo {
	serial := common.GetSerialNumberForAccessoryName(name, storage)
	return BridgeInfo{
		SerialNumber: serial,
		Password:     password,
		Name:         name,
		ID:           MAC48Address(serial),
		Manufacturer: manufacturer,
	}
}

// MAC48Address returns a MAC-48-like address from the argument string
func MAC48Address(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	result := h.Sum(nil)

	var c []string
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
