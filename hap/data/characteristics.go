package data

// Characteristics implements json of format
//
//  {
//      "characteristics": [
//          ...
//      ]
//  }
type Characteristics struct {
	Characteristics []Characteristic `json:"characteristics"`
}

// Characteristic implements json of format.
//
//  {
//      "aid": 0, "iid": 1, "value": 10 [, "status": 0, "ev": true ]
//  }
type Characteristic struct {
	AccessoryID      uint64      `json:"aid"`
	CharacteristicID uint64      `json:"iid"`
	Value            interface{} `json:"value"`

	// Status contains the status code. Should be interpreted as integer.
	// The property is omitted if not specified, which makes the payload smaller.
	Status interface{} `json:"status,omitempty"`

	// Events contains the events settings for a characteristic. Should be interpreted as boolean.
	// The property is omitted if not specified, which makes the payload smaller.
	Events interface{} `json:"ev,omitempty"`
}
