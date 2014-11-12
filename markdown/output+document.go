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
	var (
		slice = func() g.Any {
			return make([]marks, 0, 0)
		}
		list = func() g.Any {
			return g.List_.Empty()
		}

		rec func(g.List, []marks) g.Tree
	)
	rec = func(l g.List, m []marks) g.Tree {
		num := len(m)
		if num == 0 {
			return g.NewTreeNode(nothing(), l)
		}
		var (
			x        = m[num-1]
			y        = m[0 : num-1]
			nodes    = x.Children().GetOrElse(slice).([]marks)
			children = rec(g.List_.Empty(), nodes)
			node     = g.NewTreeNode(x, children.Children().GetOrElse(list).(g.List))
		)
		return rec(g.NewCons(node, l), y)
	}
	return Document{
		List: rec(g.List_.Empty(), m),
	}
}
