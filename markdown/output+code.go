package markdown

import "fmt"

type codeType string

var (
	Inline    codeType = "`"
	Multiline codeType = "```"
)

func (c codeType) String() string {
	return string(c)
}

type code struct {
	Type  codeType
	Value string
}

func (c code) String() string {
	t := c.Type.String()
	switch c.Type {
	case Inline:
		return fmt.Sprintf("%s%s%s", t, c.Value, t)
	case Multiline:
		return fmt.Sprintf("%s\n%s\n%s\n", t, c.Value, t)
	}
	return ""
}

func inline(value string) code {
	return code{
		Type:  Inline,
		Value: value,
	}
}

func multiline(value string) code {
	return code{
		Type:  Multiline,
		Value: value,
	}
}
