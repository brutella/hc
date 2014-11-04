package controller

import(
    "net/url"
    "fmt"
)

func GetCharacteristicValues(accessoryId, characteristicId int) url.Values {
    values := url.Values{}
    values.Set("id", fmt.Sprintf("%d.%d", accessoryId, characteristicId))
    
    return values
}