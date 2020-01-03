package util

import (
	"strconv"
	"strings"
)

var (
	base36 = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z"}
)

type SetupFlag uint8

var SetupFlagNone SetupFlag = 0
var SetupFlagNFC SetupFlag = 1
var SetupFlagIP SetupFlag = 2
var SetupFlagBTLE SetupFlag = 4
var SetupFlagIPWAC SetupFlag = 8

func XHMURI(pincode, setupId string, categoryId uint8, flags []SetupFlag) (string, error) {
	var version uint64 = 0
	var reserved uint64 = 0

	pincode = strings.Replace(pincode, "-", "", -1)
	code, err := strconv.ParseUint(pincode, 10, 64)
	if err != nil {
		return "", err
	}

	// Merge our flags into one
	var mergedFlags uint64
	for _, item := range flags {
		mergedFlags = mergedFlags | uint64(item)
	}

	// Build the payload
	// Payload description 45 bits (ported from https://github.com/maximkulkin/esp-homekit)
	// V = Version - 3 bits
	// R = Reserved - 4 bits
	// C = Category - 8 bits
	// F = Flags - 4 bits
	// P = Pin - 26 bits
	// VVVRRRRCCCCCCCCFFFFPPPPPPPPPPPPPPPPPPPPPPPPPP
	var payload uint64

	payload = payload | (version & 0x7)

	payload = payload << 4
	payload = payload | (reserved & 0xf)

	payload = payload << 8
	payload = payload | uint64(categoryId)

	payload = payload << 4
	payload = payload | mergedFlags&0xf

	payload = payload << 27
	payload = payload | (code & 0x7ffffff)

	setup_payload := make([]string, 9)
	for ii := 0; ii < 9; ii++ {
		setup_payload[8-ii] = base36[payload%36]
		payload = payload / 36
	}

	return "X-HM://" + strings.Join(setup_payload, "") + setupId, nil
}
