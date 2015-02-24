package controller

import (
	"fmt"
	"net/url"
)

func getCharacteristicValues(accessoryId, characteristicId int64) url.Values {
	values := url.Values{}
	values.Set("id", fmt.Sprintf("%d.%d", accessoryId, characteristicId))

	return values
}
