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

func (c codeType) String() string {
	return string(c)
}

type code struct {
	Type  codeType
	Value marks
}

func (c code) Children() g.Option {
	return g.Option_.Of(g.List_.Of(c.Value))
}

func (c code) String() string {
	t := c.Type.String()
	switch c.Type {
	case Inline:
		return fmt.Sprintf("%s%s%s", t, c.Value.String(), t)
	case Multiline:
		return fmt.Sprintf("%s\n%s\n%s\n", t, c.Value.String(), t)
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
