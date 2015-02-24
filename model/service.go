package model

// A Service is identifiable and consists of characteristics.
type Service interface {
	Compareable

	// SetId sets the service's id
	SetId(int64)

	// GetId returns the service's id
	GetId() int64

	// GetCharacteristics returns the containing characteristics
	GetCharacteristics() []Characteristic
}
