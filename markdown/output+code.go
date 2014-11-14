package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

var (
	uid int = 0
)

type codeType struct {
	value string
}

func newCodeType(val string) *codeType {
	return &codeType{
		value: val,
	}
}

var (
	Inline         *codeType = newCodeType("`")
	Multiline      *codeType = newCodeType("```")
	MultilineOpen  *codeType = newCodeType("```")
	MultilineClose *codeType = newCodeType("```")
)

func (c *codeType) IsBlock() bool {
	return false
}

func (c *codeType) Children() g.Option {
	return g.Option_.Empty()
}

func (c *codeType) String() string {
	switch c {
	case Inline:
		return c.value
	case MultilineOpen:
		return fmt.Sprintf("%s\n", c.value)
	case MultilineClose:
		return fmt.Sprintf("\n%s", c.value)
	}
	return DefaultString
}

type code struct {
	class  *codeType
	values g.List
}

func (c code) IsBlock() bool {
	return c.class != Inline
}

func (c code) Children() g.Option {
	return g.Option_.Of(c.values)
}

func (c code) String() string {
	return ""
}

func inline(val mark) code {
	return code{
		class:  Inline,
		values: g.List_.To(Inline, val, Inline),
	}
}

func multiline(val mark) code {
	return code{
		class:  Multiline,
		values: g.List_.To(MultilineOpen, val, MultilineClose),
	}
}
