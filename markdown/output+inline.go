package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type inlineType struct {
	value string
}

func newInlineType(val string) *inlineType {
	return &inlineType{
		value: val,
	}
}

var (
	Code *inlineType = newInlineType("`")
)

func (b *inlineType) IsBlock() bool {
	return false
}

func (b *inlineType) Children() g.Option {
	return g.Option_.Empty()
}

func (b *inlineType) String() string {
	return b.value
}

type inline struct {
	empty bool
	nodes g.List
	value string
}

func (b inline) IsBlock() bool {
	return false
}

func (b inline) Children() g.Option {
	if b.empty {
		return g.Option_.Empty()
	}
	return g.Option_.Of(b.nodes)
}

func (b inline) String() string {
	return b.value
}

func code(val mark) inline {
	return inline{
		nodes: g.List_.To(Code, val, Code),
	}
}

func group(val ...mark) inline {
	return inline{
		nodes: fromMarks(val),
	}
}

func link(name, url string) inline {
	return inline{
		empty: true,
		value: fmt.Sprintf("[%s](%s)", name, url),
	}
}

func nothing() inline {
	return inline{
		empty: true,
		nodes: g.List_.Empty(),
	}
}

func str(val string) inline {
	return inline{
		empty: true,
		value: val,
	}
}
