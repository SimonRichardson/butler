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

func (t unorderedListType) IsBlock() bool {
	return false
}

func (t unorderedListType) Children() g.Option {
	return g.Option_.Empty()
}

func (t unorderedListType) String() string {
	return fmt.Sprintf("%s ", string(t))
}

type orderedListType string

var (
	Hash orderedListType = "#"
)

func (t orderedListType) IsBlock() bool {
	return false
}

func (t orderedListType) Children() g.Option {
	return g.Option_.Empty()
}

func (t orderedListType) String() string {
	return fmt.Sprintf("%s ", string(t))
}

type list struct {
	nodes g.List
}

func (l list) IsBlock() bool {
	return true
}

func (l list) Children() g.Option {
	return g.Option_.Of(l.nodes)
}

func (l list) String() string {
	return ""
}

func ul(values ...mark) list {
	return list{
		nodes: fromMarks(append([]mark{Hyphen}, values...)),
	}
}

func ol(values ...mark) list {
	return list{
		nodes: fromMarks(append([]mark{Hyphen}, values...)),
	}
}
