package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type codeType string

var (
	Inline    codeType = "`"
	Multiline codeType = "```"
)

func (c codeType) Children() g.Option {
	return g.Option_.Empty()
}

func (c codeType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(c))
}

type code struct {
	Type  codeType
	Value marks
}

func (c code) Children() g.Option {
	return g.Option_.Of([]marks{c.Value})
}

func (c code) String(indent string) string {
	t := c.Type.String(DefaultIndent)
	switch c.Type {
	case Inline:
		return fmt.Sprintf("%s%s%s", t, c.Value.String(DefaultIndent), t)
	case Multiline:
		return fmt.Sprintf("%s\n%s\n%s\n", t, c.Value.String(indent), t)
	}
	return DefaultString
}

func inline(val marks) code {
	return code{
		Type:  Inline,
		Value: val,
	}
}

func multiline(val marks) code {
	return code{
		Type:  Multiline,
		Value: val,
	}
}
