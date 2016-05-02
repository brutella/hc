package swift

import (
	"bytes"
	"fmt"
	"github.com/brutella/hc/gen"
	"text/template"
)

func ServiceEnumDecl(metadata gen.Metadata) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	enum := &Enum{
		Name:  "ServiceType",
		Super: "String",
		Cases: []*Case{},
	}

	for _, svc := range metadata.Services {
		c := &Case{
			Name:  camelCased(svc.Name),
			Value: fmt.Sprintf(`"%s"`, svc.UUID),
		}
		enum.Cases = append(enum.Cases, c)
	}

	t := template.New("Test Template")

	t, err = t.Parse(EnumTemplate)
	t.Execute(&buf, enum)

	return buf.Bytes(), err
}
