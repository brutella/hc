package controller

import (
	"fmt"
	"net/url"
)

func getCharacteristicValues(accessoryID, characteristicID int64) url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d.%d", accessoryID, characteristicID))

	return values
}
