package netio

import(
    "github.com/brutella/hap"
    "os"
    "encoding/hex"
    "crypto/md5"
)

type BridgeInfo struct {
    SerialNumber string
    Password string
    Name string
    Id string
    Manufacturer string
}

func NewBridgeInfo(name, password, manufacturer string, storage hap.Storage) BridgeInfo {
    serial := hap.GetSerialNumberForAccessoryName(name, storage)
    
    return BridgeInfo{
        SerialNumber: serial,
        Password: password,
        Name: name,
        Id: IEEE802Id(serial),
        Manufacturer: manufacturer,
    }
}

// Returns the bridge id as MAC-48 address
// Is used as TXT record `id`
func IEEE802Id(input string) string {
    h := md5.New()
    h.Write([]byte(input))
    result := h.Sum(nil)
    
    bytes := result[:6]
    str := ""
    for i, b := range bytes {
        if i > 0 {
            str += ":"
        }
        str += hex.EncodeToString([]byte{b})
    }
    return str
}

// Returns the bridge name which is displayed in the accessory browser in HomeKit
// same as TXT record `md`
func Hostname() string {
    name, err := os.Hostname()
    
    if err != nil {
        name = "Unnamed"
    }
    
    return name
}