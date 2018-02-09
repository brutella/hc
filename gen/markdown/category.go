package markdown

import (
	"bytes"
	"github.com/brutella/hc/gen"
	"text/template"
)

// CatsStructTemplate is template for a CategoryMetadata struct.
const CategoriesTemplate = `| Accessory | Category |
| --- | --- |{{range .Categories}}
| {{.Name}} | {{ .Category}} | {{end}}`

// CategoriesGoCode returns the go code for a categories file
func CategoriesCode(m *gen.Metadata) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	t := template.New("Test Template")
	t, err = t.Parse(CategoriesTemplate)
	t.Execute(&buf, m)

	return buf.Bytes(), err
}
