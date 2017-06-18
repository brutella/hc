package gen

// Metadata represents the data in a HomeKit metadata file
type Metadata struct {
	Categories      []*CategoryMetadata
	Characteristics []*CharacteristicMetadata
	Services        []*ServiceMetadata
}

// CharacteristicMetadata represents a characteristic metadata entry
type CharacteristicMetadata struct {
	Constraints interface{} `json:Constraints,omitempty`
	Format      string
	Name        string
	Permissions []string
	Properties  []string `json:Properties,omitempty`
	UUID        string
	Unit        string `json:Unit,omitempty`
}

// ServiceMetadata represents a service metadata entry
type ServiceMetadata struct {
	RequiredCharacteristics []string
	OptionalCharacteristics []string
	Name                    string
	UUID                    string
}

// CategoryMetadata represents an accessory category metadata entry
type CategoryMetadata struct {
	Name     string
	Category int
}
