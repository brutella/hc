package golang

import (
	"bytes"
	"github.com/brutella/hc/gen"
	"text/template"
)

// CatsStructTemplate is template for a categories struct.
const CatsStructTemplate = `// THIS FILE IS AUTO-GENERATED
package accessory

type AccessoryType int

const ({{range .Consts}}
    {{.Identifier}} {{.TypeName}} = {{.Value}}{{end}}
)`

type Categories struct {
	Consts []*Category
}

type Category struct {
	Identifier string
	TypeName   string
	Value      int
}

func NewCategories(cats []*gen.CategoryMetadata) *Categories {
	var categories []*Category
	for _, cat := range cats {
		categories = append(categories, NewCategory(cat))
	}

	return &Categories{
		Consts: categories,
	}
}

func NewCategory(cat *gen.CategoryMetadata) *Category {
	return &Category{
		Identifier: identifier(cat),
		TypeName:   "AccessoryType",
		Value:      cat.Category,
	}
}

// CategoriesGoCode returns the go code for a categories file
func CategoriesGoCode(cats []*gen.CategoryMetadata) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	data := NewCategories(cats)

	t := template.New("Test Template")

	t, err = t.Parse(CatsStructTemplate)
	t.Execute(&buf, data)

	return buf.Bytes(), err
}

// Return the name of the characteristic type name
func identifier(cat *gen.CategoryMetadata) string {
	return "Type" + camelCased(cat.Name)
}
