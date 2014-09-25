package model

type Characteristic struct {
    Id int                  `json:"iid"` // managed by accessory
    Type CharType           `json:"type"`
    Permissions []string    `json:"perms"`
    Description string      `json:"description,omitempty"`// manufacturer description (optional) 
    
    Value interface{} `json:"value"` // any
    Format string   `json:"format"`
    Unit string     `json:"unit,omitempty"`
    
    MaxLen int          `json:"maxLen,omitempty"`
    MaxValue interface{}    `json:"maxValue,omitempty"`
    MinValue interface{}    `json:"minValue,omitempty"`
    MinStep interface{}     `json:"minStep,omitempty"`
    
    // unused
    Events bool     `json:",omitempty"`
    Bonjour bool    `json:",omitempty"`
}

// Creates a new characteristic
// If no permissions are specified, read and write will be added
func NewCharacteristic(value interface{}, format string, t CharType,  permissions []string) Characteristic {
    if len(permissions) == 0 {
        permissions = append(permissions, []string{PermRead, PermWrite}...)
    }
    
    return Characteristic{
        Value: value,
        Format: format,
        Type: t,
        Permissions: permissions,
    }
}