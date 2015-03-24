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

// NewCharacteristics returns a new characteristic.
func NewCharacteristics() *Characteristics {
	return &Characteristics{
		Characteristics: make([]Characteristic, 0),
	}
}

// AddCharacteristic adds a new characteristic.
func (r *Characteristics) AddCharacteristic(c Characteristic) {
	r.Characteristics = append(r.Characteristics, c)
}
