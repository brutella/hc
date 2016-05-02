package swift

const EnumTemplate = `enum {{.Name}}{{if .HasSuper}}: {{.Super}}{{end}} {
    {{range .Cases}}case {{.Name}} = {{.Value}}{{if .HasComment}} // {{.Comment}}{{end}}
    {{end}}
}`

type Enum struct {
	Name  string
	Super string
	Cases []*Case
}

func (e *Enum) HasSuper() bool {
	return len(e.Super) > 0
}

type Case struct {
	Name    string
	Value   string
	Comment string
}

func (c *Case) HasComment() bool {
	return len(c.Comment) > 0
}
