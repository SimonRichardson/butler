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
	H1             *blockType = newBlockType("#")
	H2             *blockType = newBlockType("##")
	H3             *blockType = newBlockType("###")
	H4             *blockType = newBlockType("####")
	H5             *blockType = newBlockType("#####")
	H6             *blockType = newBlockType("######")
	HR1            *blockType = newBlockType("======")
	HR2            *blockType = newBlockType("------")
	BlockQuote     *blockType = newBlockType(">")
	BR             *blockType = newBlockType("")
	P              *blockType = newBlockType("")
	CenterOpen     *blockType = newBlockType("->")
	CenterClose    *blockType = newBlockType("<-")
	MultilineOpen  *blockType = newBlockType("```")
	MultilineClose *blockType = newBlockType("```")
	Star           *blockType = newBlockType("*")
	Hyphen         *blockType = newBlockType("-")
	Plus           *blockType = newBlockType("+")
	Ordered        *blockType = newBlockType("1.")
)

func (b *blockType) IsBlockStart() bool {
	return false
}

func (b *blockType) IsBlockFinish() bool {
	return b == MultilineClose
}

func (b *blockType) Children() g.Option {
	return g.Option_.Empty()
}

func (b *blockType) String() string {
	switch b {
	case H1, H2, H3, H4, H5, H6:
		fallthrough
	case Star, Hyphen, Plus, Ordered:
		fallthrough
	case BlockQuote:
		return fmt.Sprintf("%s ", b.value)
	case MultilineOpen:
		return fmt.Sprintf("%s\n", b.value)
	case MultilineClose:
		return fmt.Sprintf("%s", b.value)
	}
	return b.value
}

type block struct {
	nodes g.List
}

func (b block) IsBlockStart() bool {
	return true
}

func (b block) IsBlockFinish() bool {
	return true
}

func (b block) Children() g.Option {
	return g.Option_.Of(b.nodes)
}

func (b block) String() string {
	return ""
}

func h1(val mark) block {
	return block{
		nodes: g.List_.To(H1, val),
	}
}

func h2(val mark) block {
	return block{
		nodes: g.List_.To(H2, val),
	}
}

func h3(val mark) block {
	return block{
		nodes: g.List_.To(H3, val),
	}
}

func h4(val mark) block {
	return block{
		nodes: g.List_.To(H4, val),
	}
}

func h5(val mark) block {
	return block{
		nodes: g.List_.To(H5, val),
	}
}

func h6(val mark) block {
	return block{
		nodes: g.List_.To(H6, val),
	}
}

func hr1() block {
	return block{
		nodes: g.List_.To(HR1),
	}
}

func hr2() block {
	return block{
		nodes: g.List_.To(HR2),
	}
}

func blockquote(val mark) block {
	return block{
		nodes: g.List_.To(BlockQuote, val),
	}
}

func br() block {
	return block{
		nodes: g.List_.To(BR),
	}
}

func p(val mark) block {
	return block{
		nodes: g.List_.To(P, val),
	}
}

func center(val mark) block {
	return block{
		nodes: g.List_.To(CenterOpen, val, CenterClose),
	}
}

func multiline(val mark) block {
	return block{
		nodes: g.List_.To(MultilineOpen, val, MultilineClose),
	}
}

func ul(values ...mark) block {
	return block{
		nodes: fromMarks(append(singleton(Hyphen), values...)),
	}
}

func ul०p(values ...mark) block {
	return block{
		nodes: fromMarks(append(singleton(Plus), values...)),
	}
}

func ol(values ...mark) block {
	return block{
		nodes: fromMarks(append(singleton(Ordered), values...)),
	}
}
