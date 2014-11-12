package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type blockType string

var (
	HR1        blockType = "======"
	HR2        blockType = "------"
	BlockQuote blockType = ">"
)

func (b blockType) Children() g.Option {
	return g.Option_.Empty()
}

func (b blockType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(b))
}

type block struct {
	Type  blockType
	Value marks
}

func (b block) Children() g.Option {
	switch b.Type {
	case BlockQuote:
		return g.Option_.Of([]marks{b.Value})
	}
	return g.Option_.Empty()
}

func (b block) String(indent string) string {
	switch b.Type {
	case HR1, HR2:
		return fmt.Sprintf("%s\n\n", b.Type.String(indent))
	case BlockQuote:
		return fmt.Sprintf("%s %s\n", b.Type.String(indent), b.Value.String(DefaultIndent))
	}
	return DefaultString
}

func hr1() block {
	return block{
		Type: HR1,
	}
}

func hr2() block {
	return block{
		Type: HR2,
	}
}

func blockquote(val marks) block {
	return block{
		Type:  BlockQuote,
		Value: val,
	}
}
