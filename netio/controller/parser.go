package controller

import (
	"github.com/brutella/hap/common"
	"github.com/gosexy/to"
	"strings"
)

// ParseAccessoryAndCharacterId returns the accessory and characteristic id encoded in the argument string.
// The string must be in format "<accessory id>.<characteristic id>"
func ParseAccessoryAndCharacterId(str string) (int64, int64, error) {
	ids := strings.Split(str, ".")
	if len(ids) != 2 {
		return 0, 0, common.NewErrorf("Could not parse uid %s", str)
	}

	aid := to.Int64(ids[0])
	cid := to.Int64(ids[1])

	return aid, cid, nil
}
