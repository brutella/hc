package golang

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/brutella/hc/gen"
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

    {{range .Optional}}
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

func (svc *{{.StructName}}) AddOptionalCharacteristics() {
   {{range .Optional}}
   svc.{{.StructName}} = characteristic.New{{.StructName}}()
   svc.AddCharacteristic(svc.{{.StructName}}.Characteristic)
   {{end}}
}
`

// Service holds service template data
type Service struct {
	Name          string // Name of the service (e.g. Light bulb)
	StructName    string // Name of the struct (e.g. Lightbulb)
	FileName      string // Name of the file (e.g. lightbulb.go)
	LocalFilePath string // Path to the file (e.g. ~/User/Go/src/github.com/brutella/hc/service/lightbulb.go)
	RelFilePath   string // Relative path to the file from the project root (e.g. service/lightbulb.go)
	TypeName      string // Name of type e.g. TypeLightbulb
	TypeValue     string // Value of the type e.g. 00000008-0000-1000-8000-0026BB765291

	Chars    []*Characteristic
	Optional []*Characteristic
}

// FileName returns the filename for a characteristic
func ServiceFileName(svc *gen.ServiceMetadata) string {
	return fmt.Sprintf("%s.go", underscored(svc.Name))
}

func ServiceDecl(svc *gen.ServiceMetadata, chars []*gen.CharacteristicMetadata) *Service {
	return &Service{
		Name:          svc.Name,
		StructName:    camelCased(svc.Name),
		FileName:      ServiceFileName(svc),
		LocalFilePath: ServiceLocalFilePath(svc),
		RelFilePath:   ServiceRelativeFilePath(svc),
		TypeName:      serviceTypeName(svc),
		TypeValue:     minifyUUID(svc.UUID),
		Chars:         requiredCharacteristics(svc, chars),
		Optional:      optionalCharacteristics(svc, chars),
	}
}

// ServiceGoCode returns the o code for a characteristic file
func ServiceGoCode(svc *gen.ServiceMetadata, chars []*gen.CharacteristicMetadata) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	data := ServiceDecl(svc, chars)

	t := template.New("Test Template")

	t, err = t.Parse(ServiceStructTemplate)
	t.Execute(&buf, data)

	return buf.Bytes(), err
}

// Return the name of the characteristic type name
func serviceTypeName(svc *gen.ServiceMetadata) string {
	return "Type" + camelCased(svc.Name)
}

func requiredCharacteristics(svc *gen.ServiceMetadata, chars []*gen.CharacteristicMetadata) []*Characteristic {
	var required = []*Characteristic{}
	for _, uuid := range svc.RequiredCharacteristics {
		char := charWithUUID(uuid, chars)
		data := NewCharacteristic(char)
		required = append(required, data)
	}

	return required
}

func optionalCharacteristics(svc *gen.ServiceMetadata, chars []*gen.CharacteristicMetadata) []*Characteristic {
	var required = []*Characteristic{}
	for _, uuid := range svc.OptionalCharacteristics {
		char := charWithUUID(uuid, chars)
		data := NewCharacteristic(char)
		required = append(required, data)
	}

	return required
}

func charWithUUID(uuid string, chars []*gen.CharacteristicMetadata) *gen.CharacteristicMetadata {
	for _, char := range chars {
		if char.UUID == uuid {
			return char
		}
	}

	return nil
}
