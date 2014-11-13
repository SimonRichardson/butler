package markdown

import g "github.com/SimonRichardson/butler/generic"

type blockType string

var (
	HR1        blockType = "======"
	HR2        blockType = "------"
	BlockQuote blockType = ">"
)

func (b blockType) IsInline() bool {
	return false
}

func (b blockType) Children() g.Option {
	return g.Option_.Empty()
}

func (b blockType) String() string {
	return string(b)
}

type block struct {
	values g.List
}

func (b block) IsInline() bool {
	return false
}

func (b block) Children() g.Option {
	return g.Option_.Of(b.values)
}

func (b block) String() string {
	return ""
}

func hr1() block {
	return block{
		values: g.List_.To(HR1),
	}
}

func hr2() block {
	return block{
		values: g.List_.To(HR2),
	}
}

func blockquote(val marks) block {
	return block{
		values: g.List_.To(BlockQuote, val),
	}
}
