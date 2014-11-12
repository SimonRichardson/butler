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

func (t unorderedListType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(t))
}

type orderedListType string

var (
	Hash orderedListType = "#"
)

func (t orderedListType) String(indent string) string {
	return fmt.Sprintf("%s%s", indent, string(t))
}

type list struct {
	Type  marks
	Value marks
}

func (l list) String(indent string) string {
	return fmt.Sprintf("%s %s\n", l.Type.String(indent), l.Value.String(DefaultIndent))
}

func ul(values []marks) g.Tree {
	return make(values, func(a marks) list {
		return li(Hyphen, a)
	})
}

func ol(values []marks) g.Tree {
	return make(values, func(a marks) list {
		return li(Hash, a)
	})
}

func li(a marks, b marks) list {
	return list{
		Type:  a,
		Value: b,
	}
}

func make(values []marks, f func(marks) list) g.Tree {
	var rec func(g.List, []marks) g.List
	rec = func(a g.List, b []marks) g.List {
		num := len(b)
		if num == 0 {
			return a
		}
		var (
			x = num - 1
			y = g.Tree_.Of(f(b[x]))
			z = b[0:x]
		)
		return rec(g.NewCons(y, a), z)
	}
	return g.NewTreeNode(
		nothing(),
		rec(g.List_.Empty(), values),
	)
}
