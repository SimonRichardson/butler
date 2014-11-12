package markdown

import "fmt"

type codeType string

var (
	Inline    codeType = "`"
	Multiline codeType = "```"
)

func (c codeType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(c))
}

type code struct {
	Type  codeType
	Value raw
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

func inline(val string) code {
	return code{
		Type:  Inline,
		Value: value(val),
	}
}

func multiline(val string) code {
	return code{
		Type:  Multiline,
		Value: value(val),
	}
}
