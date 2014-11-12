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

func (t unorderedListType) Children() g.Option {
	return g.Option_.Empty()
}

func (t unorderedListType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(t))
}

type orderedListType string

var (
	Hash orderedListType = "#"
)

func (t orderedListType) Children() g.Option {
	return g.Option_.Empty()
}

func (t orderedListType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(t))
}

type list struct {
	Type  marks
	nodes []marks
}

func (l list) Children() g.Option {
	return g.Option_.Of(l.nodes)
}

func (l list) String(indent string) string {
	return fmt.Sprintf("%s\n", l.Type.String(indent))
}

func ul(values ...marks) list {
	return list{
		Type:  Hyphen,
		nodes: values,
	}
}

func ol(values ...marks) list {
	return list{
		Type:  Hash,
		nodes: values,
	}
}
