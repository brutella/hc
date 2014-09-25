package hap

type Characteristic struct {
    Id int                  `json:"iid"`
    Type UUID               `json:"type"`
    Permissions []string    `json:"perms"`
    Description string      `json:"description"`// manufacturer description (optional) 
    
    Value string    `json:"value"` // any
    Format string   `json:"format"`
    Unit string     `json:"unit"`
    
    MaxLen int          `json:"maxLen,omitempty"`
    MaxValue float64    `json:"maxValue,omitempty"`
    MinValue float64    `json:"minValue,omitempty"`
    MinStep float64     `json:"minStep,omitempty"`
    
    // unused
    Events bool     `json:",omitempty"`
    Bonjour bool    `json:",omitempty"`
}

func NewCharacteristic() *Characteristic {
    return &Characteristic{}
}