package markdown

import (
	g "github.com/SimonRichardson/butler/generic"
)

type BlockKind int

var (
	None   BlockKind = 1 << 0
	Start  BlockKind = 1 << 1
	Finish BlockKind = 1 << 2
	Before BlockKind = 1 << 3
	After  BlockKind = 1 << 4
	Pad    BlockKind = 1 << 5
)

type Markup interface {
	Concat(y Markup) Markup
}

type parent struct {
	tag, start, finish string
	value              Markup
	kind               BlockKind
}

func (m parent) Concat(y Markup) Markup {
	return Concat(m, y)
}

type leaf struct {
	tag, start, finish string
	kind               BlockKind
}

func (m leaf) Concat(y Markup) Markup {
	return Concat(m, y)
}

type content struct {
	value string
}

func (m content) Concat(y Markup) Markup {
	return Concat(m, y)
}

type concat struct {
	left, right Markup
}

func (m concat) Concat(y Markup) Markup {
	return Concat(m, y)
}

type empty struct{}

func (m empty) Concat(y Markup) Markup {
	return y
}

func Parent(tag, start, finish string, kind BlockKind) func(Markup) Markup {
	return func(value Markup) Markup {
		return parent{
			tag:    tag,
			start:  start,
			finish: finish,
			kind:   kind,
			value:  value,
		}
	}
}

func Leaf(tag, start, finish string, kind BlockKind) Markup {
	return leaf{
		tag:    tag,
		start:  start,
		finish: finish,
		kind:   kind,
	}
}

func Content(value string) Markup {
	return content{
		value: value,
	}
}

func Concat(left, right Markup) Markup {
	return concat{
		left:  left,
		right: right,
	}
}

func Empty() Markup {
	return empty{}
}

var (
	Markup_ = markup_{}
)

type markup_ struct{}

func (m markup_) Of(x g.Any) Markup {
	return Content(x.(string))
}

func (m markup_) Empty(x g.Any) Markup {
	return Empty()
}
