package output

import (
	"bytes"
	"text/template"

	"github.com/SimonRichardson/butler/generic"
)

type TextEncoder struct {
	Template string
}

func (e TextEncoder) Encode(a generic.Any) ([]byte, error) {
	var (
		buffer *bytes.Buffer
	)
	tmpl := template.Must(template.New("text-encoder").Parse(e.Template))
	if err := tmpl.Execute(buffer, a); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
