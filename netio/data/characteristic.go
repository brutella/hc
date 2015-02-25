package data

// Implements json of format
//
//  {
//      "aid": 0, "iid": 1, "value": 10 [, "ev": true ]
//  }
type Characteristic struct {
	AccessoryId int64       `json:"aid"`
	Id          int64       `json:"iid"`
	Value       interface{} `json:"value"`

	// Events property is true or false
	Events interface{} `json:"ev,omitempty"`
}
