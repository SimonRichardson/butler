package io

import (
	"bytes"
	"html/template"

	g "github.com/SimonRichardson/butler/generic"
)

type HtmlEncoder struct {
	Template string
}

func (e HtmlEncoder) Encode(a g.Any) g.Either {
	var (
		buffer *bytes.Buffer
	)
	tmpl, err := template.New("html-encoder").Parse(e.Template)
	if err != nil {
		return g.NewLeft(err)
	}
	if err := tmpl.Execute(buffer, a); err != nil {
		return g.NewLeft(err)
	}
	return g.NewRight(buffer.Bytes())
}

func (e HtmlEncoder) Keys(a g.Any) g.Either {
	return g.NewLeft(a)
}

func (e HtmlEncoder) Generate(x g.Any) g.Either {
	return generate(e)(x)
}
