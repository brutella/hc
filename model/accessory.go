package model

type Accessory struct {
    UUId int `json:"aid"`
    Services []Service `json:"services"`
    
    Characteristics []Characteristic
}

type AccessoryInfo struct {
    Accessory
}

// func NewAccessoryInfo(name string) (*AccessoryInfo, error){
//     a := Accessory{Name: name, Password: password}
//
//     return &a, err
// }