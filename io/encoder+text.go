package io

import (
	"bytes"
	"text/template"

	g "github.com/SimonRichardson/butler/generic"
)

type TextEncoder struct {
	Template string
}

func (e TextEncoder) Encode(a g.Any) g.Either {
	var (
		buffer *bytes.Buffer
	)
	tmpl := template.Must(template.New("text-encoder").Parse(e.Template))
	if err := tmpl.Execute(buffer, a); err != nil {
		return g.NewLeft(err)
	}
	return g.NewRight(buffer.Bytes())
}

func (e TextEncoder) Keys(a g.Any) g.Either {
	return g.NewLeft(a)
}

func (e TextEncoder) Generate(x g.Any) g.Either {
	return generate(e)(x)
}

func (e TextEncoder) String() string {
	return "TextEncoder"
}
