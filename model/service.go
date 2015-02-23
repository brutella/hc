package model

type Service interface {
	Compareable

	// SetId sets the service's id
	SetId(int64)

	// GetId returns the service's id
	GetId() int64

	// GetCharacteristics returns a list of characteristic which represent the service
	GetCharacteristics() []Characteristic
}
