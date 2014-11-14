package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type blockType struct {
	value string
}

func newBlockType(val string) *blockType {
	return &blockType{
		value: val,
	}
}

var (
	HR1        *blockType = newBlockType("======")
	HR2        *blockType = newBlockType("------")
	BlockQuote *blockType = newBlockType(">")
	BR         *blockType = newBlockType("")
	P          *blockType = newBlockType("")
)

func (b *blockType) IsBlock() bool {
	return false
}

func (b *blockType) Children() g.Option {
	return g.Option_.Empty()
}

func (b *blockType) String() string {
	if b == BlockQuote {
		return fmt.Sprintf("%s ", b.value)
	}
	return b.value
}

type block struct {
	values g.List
}

func (b block) IsBlock() bool {
	return true
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

func br() block {
	return block{
		values: g.List_.To(BR),
	}
}

func p(val marks) block {
	return block{
		values: g.List_.To(P, val),
	}
}
