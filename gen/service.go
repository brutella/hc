package gen

import (
	"bytes"
	"fmt"
	"text/template"
)

// ServiceStructTemplate is template for a service struct.
const ServiceStructTemplate = `// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const {{.TypeName}} = "{{.TypeValue}}"

type {{.StructName}} struct {
    *Service
    {{range .Chars}}
    {{.StructName}} *characteristic.{{.StructName}}{{end}}
}

func New{{.StructName}}() *{{.StructName}} {
    svc := {{.StructName}}{}
    svc.Service = New({{.TypeName}})
    {{range .Chars}}
    svc.{{.StructName}} = characteristic.New{{.StructName}}()
    svc.AddCharacteristic(svc.{{.StructName}}.Characteristic)
    {{end}}
    
	return &svc
}
`

// Service holds service template data
type Service struct {
	StructName string // Name of the struct (e.g. Lightbulb)
	FileName   string // Name of the file (e.g. lightbulb.go)
	TypeName   string // Name of type e.g. TypeLightbulb
	TypeValue  string // Value of the type e.g. 00000008-0000-1000-8000-0026BB765291

	Chars []*Characteristic
}

// FileName returns the filename for a characteristic
func ServiceFileName(svc *ServiceMetadata) string {
	return fmt.Sprintf("%s.go", underscored(svc.Name))
}

// ServiceGoCode returns the o code for a characteristic file
func ServiceGoCode(svc *ServiceMetadata, chars []*CharacteristicMetadata) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	data := Service{
		StructName: camelCased(svc.Name),
		FileName:   ServiceFileName(svc),
		TypeName:   serviceTypeName(svc),
		TypeValue:  svc.UUID,
		Chars:      requiredCharacteristics(svc, chars),
	}

	t := template.New("Test Template")

	t, err = t.Parse(ServiceStructTemplate)
	t.Execute(&buf, data)

	return buf.Bytes(), err
}

// Return the name of the characteristic type name
func serviceTypeName(svc *ServiceMetadata) string {
	return "Type" + camelCased(svc.Name)
}

func requiredCharacteristics(svc *ServiceMetadata, chars []*CharacteristicMetadata) []*Characteristic {
	var required = []*Characteristic{}
	for _, uuid := range svc.RequiredCharacteristics {
		char := charWithUUID(uuid, chars)
		data := NewCharacteristic(char)
		required = append(required, data)
	}

	return required
}

func charWithUUID(uuid string, chars []*CharacteristicMetadata) *CharacteristicMetadata {
	for _, char := range chars {
		if char.UUID == uuid {
			return char
		}
	}

	return nil
}
