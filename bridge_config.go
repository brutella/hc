package hap

import(
    "os"
    "crypto/md5"
    "encoding/hex"
)

type BridgeInfo struct {
    SerialNumber string
    Password string
    Name string
    Id string
    Manufacturer string
}

func NewBridgeInfo(name, password, serialNumber, manufacturer string) BridgeInfo {
    return BridgeInfo{
        SerialNumber: serialNumber,
        Password: password,
        Name: name,
        Id: IEEE802Id(),
        Manufacturer: manufacturer,
    }
}

// Returns the bridge id as MAC-48 address
// Is used as TXT record `id`
func IEEE802Id() string {
    h := md5.New()
    h.Write([]byte(Hostname()))
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