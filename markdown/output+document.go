package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type Document struct {
	List g.Tree
}

func (d Document) Children() g.Option {
	return g.Option_.Empty()
}

func (d Document) String() string {
	return d.List.FoldLeft("", func(a, b g.Any) g.Any {
		return fmt.Sprintf("%s%s", a.(string), b.(marks).String())
	}).(string)
}

func document(m ...marks) Document {
	list := func() g.Any {
		return g.List_.Empty()
	}

	var rec func(g.List, g.List) g.Tree
	rec = func(l g.List, m g.List) g.Tree {
		return m.Head().Fold(
			func(a g.Any) g.Any {
				var (
					x        = a.(marks)
					y        = m.Tail()
					nodes    = x.Children().GetOrElse(list).(g.List)
					children = rec(g.List_.Empty(), nodes)
					node     = g.NewTreeNode(x, children.Children().GetOrElse(list).(g.List))
				)
				return rec(g.NewCons(node, l), y)
			},
			func() g.Any {
				return g.NewTreeNode(nothing(), l)
			},
		).(g.Tree)
	}
	return Document{
		List: rec(g.List_.Empty(), fromMarks(m)),
	}
}
