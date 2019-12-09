package markdown

import (
	"bytes"
	"github.com/brutella/hc/gen"
	"github.com/brutella/hc/gen/golang"
	"text/template"
)

const ServicesTemplate = `| Service | Characteristics | ID
| --- | --- | --- |
{{range .Svcs}}| <a href="../{{.RelFilePath}}">{{.Name}}</a> | {{range $idx, $ch := .Chars}}{{ if $idx }}<br/>{{ end }}<a href="../{{$ch.RelFilePath}}">{{ $ch.Name }}</a>{{end}}{{range $idx, $ch := .Optional}}<br/><a href="../{{$ch.RelFilePath}}">{{ $ch.Name }}</a> <small>Optional</small>{{end}} | {{ .TypeValue }} |
{{end}}`

type Services struct {
	Svcs []*golang.Service
}

func ServicesCode(m *gen.Metadata) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	var svcs []*golang.Service
	for _, svc := range m.Services {
		svcs = append(svcs, golang.ServiceDecl(svc, m.Characteristics))
	}

	data := Services{svcs}
	t := template.New("Test Template")
	t, err = t.Parse(ServicesTemplate)
	t.Execute(&buf, data)

	return buf.Bytes(), err
}
