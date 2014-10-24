package output

import (
	"bytes"
	"html/template"

	"github.com/SimonRichardson/butler/generic"
)

type HtmlEncoder struct {
	Template string
}

func (e HtmlEncoder) Encode(a generic.Any) ([]byte, error) {
	var (
		buffer *bytes.Buffer
	)
	tmpl, err := template.New("html-encoder").Parse(e.Template)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buffer, a); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
