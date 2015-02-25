package data

// Implements json of format
//
//  {
//      "characteristics": [
//          ...
//      ]
//  }
type Characteristics struct {
	Characteristics []Characteristic `json:"characteristics"`
}

func NewCharacteristics() *Characteristics {
	return &Characteristics{
		Characteristics: make([]Characteristic, 0),
	}
}

func (r *Characteristics) AddCharacteristic(c Characteristic) {
	r.Characteristics = append(r.Characteristics, c)
}
