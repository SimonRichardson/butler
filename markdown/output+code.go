package markdown

import g "github.com/SimonRichardson/butler/generic"

type codeType string

var (
	Inline    codeType = "`"
	Multiline codeType = "```"
)

func (c codeType) IsInline() bool {
	return c == Inline
}

func (c codeType) Children() g.Option {
	return g.Option_.Empty()
}

func (c codeType) String() string {
	return string(c)
}

type code struct {
	values g.List
}

func (c code) IsInline() bool {
	return false
}

func (c code) Children() g.Option {
	return g.Option_.Of(c.values)
}

func (c code) String() string {
	return ""
}

func inline(val marks) code {
	return code{
		values: g.List_.To(Inline, val, Inline),
	}
}

func multiline(val marks) code {
	return code{
		values: g.List_.To(Multiline, val, Multiline),
	}
}
