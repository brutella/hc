package gen

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

// CharStructTemplate is template for a characteristic struct.
const CharStructTemplate = `// THIS FILE IS AUTO-GENERATED
package characteristic

{{if .HasConsts}}
const (
    {{range .Consts}}
    {{.Identifier}} {{.TypeName}} = {{.Value}}{{end}}
)
{{end}}

const {{.TypeName}} = "{{.TypeValue}}"

type {{.StructName}} struct {
    *{{.EmbeddedStructName}}
}

func New{{.StructName}}() *{{.StructName}} {
    char := New{{.EmbeddedStructName}}({{.TypeName}})
    char.Format = {{.FormatTypeName}}
    char.Perms = {{.PermsDecl}}
    {{if .HasMinValue}}char.SetMinValue({{.MinValue}}){{end}}
    {{if .HasMaxValue}}char.SetMaxValue({{.MaxValue}}){{end}}
    {{if .HasStepValue}}char.SetStepValue({{.StepValue}}){{end}}
    {{if .HasDefaultValue}}char.SetValue({{.DefaultValue}}){{end}}
    {{if .UnitName}}char.Unit = {{.UnitName}}{{end}}
    
	return &{{.StructName}}{char}
}`

// Characteristic holds characteristic template data
type Characteristic struct {
	EmbeddedStructName string      // Name of the embedded struct (e.g. Int)
	FormatTypeName     string      // Name of the format type (e.g. FormatInt32)
	StructName         string      // Name of the struct (e.g. Brightness)
	FileName           string      // Name of the file (e.g. brightness.go)
	PermsDecl          string      // Permissions declaration (e.g. []string{PermRead, PermWrite, PermEvents})
	TypeName           string      // Name of type e.g. TypeBrightness
	TypeValue          string      // Value of the type e.g. 00000008-0000-1000-8000-0026BB765291
	DefaultValue       interface{} // e.g. 0
	MinValue           interface{} // e.g. 0
	MaxValue           interface{} // e.g. 100
	StepValue          interface{} // e.g. 1
	UnitName           string      // Name of the unit e.g. UnitPercentage

	Consts []ConstDecl
}

func NewCharacteristic(char *CharacteristicMetadata) *Characteristic {
	data := Characteristic{
		EmbeddedStructName: embeddedStructNames[char.Format],
		FormatTypeName:     formatConstants[char.Format],
		StructName:         structName(char),
		FileName:           FileName(char),
		PermsDecl:          permissionDecl(char),
		TypeName:           typeName(char),
		TypeValue:          minifyUUID(char.UUID),
		DefaultValue:       defaultValue(char),
		MinValue:           minValue(char),
		MaxValue:           maxValue(char),
		StepValue:          stepValue(char),
		UnitName:           unitName(char),
		Consts:             constDecls(char),
	}

	return &data
}

// HasMinValue returns true if characteristic has a min value
func (d Characteristic) HasMinValue() bool {
	return d.MinValue != nil
}

// HasMaxValue returns true if characteristic has a max value
func (d Characteristic) HasMaxValue() bool {
	return d.MaxValue != nil
}

// HasStepValue returns true if characteristic has a step value
func (d Characteristic) HasStepValue() bool {
	return d.StepValue != nil
}

// HasConsts returns true if characteristic has const declarations
func (d Characteristic) HasConsts() bool {
	return len(d.Consts) > 0
}

// HasDefaultValue returns true if characteristic has a default value
func (d Characteristic) HasDefaultValue() bool {
	return d.DefaultValue != nil
}

// ConstDecl is a constant declaration
type ConstDecl struct {
	Identifier string
	TypeName   string
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

// FileName returns the filename for a characteristic
func FileName(char *CharacteristicMetadata) string {
	return fmt.Sprintf("%s.go", underscored(char.Name))
}

// CharacteristicGoCode returns the o code for a characteristic file
func CharacteristicGoCode(char *CharacteristicMetadata) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	data := NewCharacteristic(char)

	t := template.New("Test Template")

	t, err = t.Parse(CharStructTemplate)
	t.Execute(&buf, data)

	return buf.Bytes(), err
}

var formatConstants = map[string]string{
	"string": "FormatString",
	"bool":   "FormatBool",
	"float":  "FormatFloat",
	"uint8":  "FormatUInt8",
	"uint16": "FormatUInt16",
	"uint32": "FormatUInt32",
	"int32":  "FormatInt32",
	"uint64": "FormatUInt64",
	"tlv8":   "FormatTLV8",
}

var constTypes = map[string]string{
	"string": "string",
	"bool":   "bool",
	"float":  "float",
	"uint8":  "int",
	"uint16": "int",
	"uint32": "int",
	"int32":  "int",
	"uint64": "int",
	"tlv8":   "[]byte",
}

var embeddedStructNames = map[string]string{
	"string": "String",
	"bool":   "Bool",
	"float":  "Float",
	"uint8":  "Int",
	"uint16": "Int",
	"uint32": "Int",
	"int32":  "Int",
	"tlv8":   "Bytes",
}

// isReadable returns true the characteristic contains the readable property
func isReadable(char *CharacteristicMetadata) bool {
	for _, perm := range char.Properties {
		if perm == "read" {
			return true
		}
	}

	return false
}

// defaultValue returns the default value of a characteristic, based on the characteristic format and properties (readable)
func defaultValue(char *CharacteristicMetadata) interface{} {
	if isReadable(char) == false {
		return nil
	}

	switch char.Format {
	case "string":
		return `""`
	case "bool":
		return "false"
	case "int", "float", "uint8", "uint16", "uint32", "int32", "uint64":
		if min := minValue(char); min != nil {
			return min
		}

		return 0
	case "tlv8":
		return "[]byte{}"
	default:
		break
	}

	return nil
}

// minifyUUID returns a minified version of s by removing unneeded characters.
// For example the UUID "0000008C-0000-1000-8000-0026BB765291" the Window Covering
// service will be minified to "8C".
func minifyUUID(s string) string {
	authRegexp := regexp.MustCompile(`^([0-9a-fA-F]*)`)
	if str := authRegexp.FindString(s); len(str) > 0 {
		return strings.TrimLeft(str, "0")
	}

	return s
}

// Return the name of the characteristic type name
func typeName(char *CharacteristicMetadata) string {
	return "Type" + camelCased(char.Name)
}

func structName(char *CharacteristicMetadata) string {
	return camelCased(char.Name)
}

// strip removes any leading and trailing white spaces, and make the following substitutions: "." => "_", ","|"-" => "" (empty string)
func strip(s string) string {
	trimmed := strings.TrimSpace(s)

	r := strings.NewReplacer(".", "_", ",", "", "-", "")
	return r.Replace(trimmed)
}

func underscored(s string) string {
	lowered := strings.ToLower(strip(s))
	return strings.Replace(lowered, " ", "_", -1)
}

func camelCased(s string) string {
	lowered := strings.Title(strip(s))
	return strings.Replace(lowered, " ", "", -1)
}

func permissionDecl(char *CharacteristicMetadata) string {
	var perms []string
	for _, perm := range char.Properties {
		switch perm {
		case "read":
			perms = append(perms, "PermRead")
		case "write":
			perms = append(perms, "PermWrite")
		case "cnotify":
			perms = append(perms, "PermEvents")
		case "uncnotify":
			// TODO(mah)
			break
		default:
			log.Fatal(fmt.Sprintf("Undefined characteristic permission %s", perm))
		}
	}

	return "[]string{" + strings.Join(perms, ",") + "}"
}

func unitName(char *CharacteristicMetadata) string {
	switch char.Unit {
	case "percentage":
		return "UnitPercentage"
	case "arcdegrees":
		return "UnitArcDegrees"
	case "celsius":
		return "UnitCelsius"
	default:
		return ""
	}
}

func constraints(char *CharacteristicMetadata) map[string]interface{} {
	if char.Constraints != nil {
		if constraints, ok := char.Constraints.(map[string]interface{}); ok == true {
			return constraints
		}
	}
	return nil
}

func constraintWithKey(char *CharacteristicMetadata, key string) interface{} {
	if constr := constraints(char); constr != nil {
		return constr[key]
	}
	return nil
}

func constrainedValues(char *CharacteristicMetadata) map[string]interface{} {
	if values := constraintWithKey(char, "ValidValues"); values != nil {
		return values.(map[string]interface{})
	}

	return nil
}

func minValue(char *CharacteristicMetadata) interface{} {
	return constraintWithKey(char, "MinimumValue")
}

func maxValue(char *CharacteristicMetadata) interface{} {
	return constraintWithKey(char, "MaximumValue")
}

func stepValue(char *CharacteristicMetadata) interface{} {
	return constraintWithKey(char, "StepValue")
}

func constDecls(char *CharacteristicMetadata) []ConstDecl {
	if values := constrainedValues(char); values != nil {
		name := camelCased(char.Name)
		typ := constTypes[char.Format]

		var decls []ConstDecl
		for key, value := range values {
			str := fmt.Sprintf("%s", value)

			c := ConstDecl{name + camelCased(str), typ, key}
			decls = append(decls, c)
		}

		// Sort by value in ascending order
		sort.Sort(ByValue(decls))
		return decls
	}

	return nil
}
