package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type unorderedListType string

var (
	Star   unorderedListType = "*"
	Hyphen unorderedListType = "-"
)

func (t unorderedListType) IsInline() bool {
	return false
}

func (t unorderedListType) Children() g.Option {
	return g.Option_.Empty()
}

func (t unorderedListType) String() string {
	return string(t)
}

type orderedListType string

var (
	Hash orderedListType = "#"
)

func (t orderedListType) IsInline() bool {
	return false
}

func (t orderedListType) Children() g.Option {
	return g.Option_.Empty()
}

func (t orderedListType) String() string {
	return string(t)
}

type list struct {
	Type  marks
	nodes g.List
}

func (l list) IsInline() bool {
	return false
}

func (l list) Children() g.Option {
	return g.Option_.Of(l.nodes)
}

func (l list) String() string {
	return fmt.Sprintf("%s", l.Type.String())
}

func ul(values ...marks) list {
	return list{
		Type:  Hyphen,
		nodes: fromMarks(values),
	}
}

func ol(values ...marks) list {
	return list{
		Type:  Hash,
		nodes: fromMarks(values),
	}
}
