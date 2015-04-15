package controller

import (
	"fmt"
	"github.com/gosexy/to"
	"strings"
)

// ParseAccessoryAndCharacterID returns the accessory and characteristic id encoded in the argument string.
// The string must be in format "<accessory id>.<characteristic id>"
func ParseAccessoryAndCharacterID(str string) (accessoryID int64, characteristicID int64, err error) {
	ids := strings.Split(str, ".")
	if len(ids) != 2 {
		err = fmt.Errorf("Could not parse uid %s", str)
	} else {
		accessoryID = to.Int64(ids[0])
		characteristicID = to.Int64(ids[1])
	}

	return accessoryID, characteristicID, err
}
