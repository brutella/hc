package model

// A Service is identifiable and consists of characteristics.
type Service interface {
	Compareable

	// SetID sets the service's id
	SetID(int64)

	// GetID returns the service's id
	GetID() int64

	// GetCharacteristics returns the containing characteristics
	GetCharacteristics() []Characteristic
}
