package swift

import (
	"bytes"
	"fmt"
	"github.com/brutella/hc/gen"
	"log"
	"reflect"
	"sort"
	"strings"
	"text/template"
)

// ConstDecl is a constant declaration
type ConstDecl struct {
	Identifier string
	Value      interface{}
}

// ByValue defines a type for sorting const declarations by value by Impelementing the sort.Interface interface
type ByValue []ConstDecl

func (v ByValue) Len() int {
	return len(v)
}

func (v ByValue) Less(i, j int) bool {
	iValue, iOK := v[i].Value.(string)
	jValue, jOK := v[j].Value.(string)

	if iOK == false {
		log.Fatalf("Expected string, got %v", reflect.TypeOf(v[i].Value))
	}
	if jOK == false {
		log.Fatalf("Expected string, got %v", reflect.TypeOf(v[j].Value))
	}

	return iValue < jValue
}

func (v ByValue) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func CharacteristicEnumDecl(metadata gen.Metadata) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	enum := &Enum{
		Name:  "CharacteristicType",
		Super: "String",
		Cases: []*Case{},
	}

	for _, char := range metadata.Characteristics {
		comment := char.Format
		consts := []string{}
		for _, c := range constDecls(char) {
			consts = append(consts, fmt.Sprintf("%v: %s", c.Value, c.Identifier))
		}

		if len(consts) > 0 {
			comment = fmt.Sprintf("%s (%s)", comment, strings.Join(consts, ", "))
		}

		c := &Case{
			Name:    typeName(char),
			Value:   fmt.Sprintf(`"%s"`, char.UUID),
			Comment: comment,
		}
		enum.Cases = append(enum.Cases, c)
	}

	t := template.New("Test Template")

	t, err = t.Parse(EnumTemplate)
	t.Execute(&buf, enum)

	return buf.Bytes(), err
}

// Return the name of the characteristic type name
func typeName(char *gen.CharacteristicMetadata) string {
	return camelCased(char.Name)
}

// strip removes any leading and trailing white spaces, and make the following substitutions: "." => "_", ","|"-" => "" (empty string)
func strip(s string) string {
	trimmed := strings.TrimSpace(s)

	r := strings.NewReplacer(".", "_", ",", "", "-", "")
	return r.Replace(trimmed)
}

func camelCased(s string) string {
	lowered := strings.Title(strip(s))
	return strings.Replace(lowered, " ", "", -1)
}

func constraints(char *gen.CharacteristicMetadata) map[string]interface{} {
	if char.Constraints != nil {
		if constraints, ok := char.Constraints.(map[string]interface{}); ok == true {
			return constraints
		}
	}
	return nil
}

func constraintWithKey(char *gen.CharacteristicMetadata, key string) interface{} {
	if constr := constraints(char); constr != nil {
		return constr[key]
	}
	return nil
}

func constrainedValues(char *gen.CharacteristicMetadata) map[string]interface{} {
	if values := constraintWithKey(char, "ValidValues"); values != nil {
		return values.(map[string]interface{})
	}

	return nil
}

func constDecls(char *gen.CharacteristicMetadata) []ConstDecl {
	if values := constrainedValues(char); values != nil {

		var decls []ConstDecl
		for key, value := range values {
			str := fmt.Sprintf("%s", value)

			c := ConstDecl{str, key}
			decls = append(decls, c)
		}

		// Sort by value in ascending order
		sort.Sort(ByValue(decls))
		return decls
	}

	return nil
}
