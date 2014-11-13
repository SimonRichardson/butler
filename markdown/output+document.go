package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type Document struct {
	List g.Tree
}

func (d Document) IsInline() bool {
	return false
}

func (d Document) Children() g.Option {
	return g.Option_.Empty()
}

func (d Document) String() string {
	return d.List.FoldLeft("", func(a, b g.Any) g.Any {
		var (
			x = a.(string)
			y = b.(depthNode)
			z = y.node.String()
		)
		if y.node.IsInline() {
			return fmt.Sprintf("%s%s", x, z)
		} else {
			return fmt.Sprintf("%s\n%s%s", x, indent(y.depth), z)
		}
	}).(string)
}

func document(m ...marks) Document {
	var (
		empty = g.List_.Empty()
		list  = func() g.Any {
			return empty
		}
		rec func(g.List, g.List, int) g.Tree
	)
	rec = func(l g.List, m g.List, depth int) g.Tree {
		return m.Head().Fold(
			func(a g.Any) g.Any {
				var (
					x        = a.(depthNode)
					y        = m.Tail()
					nodes    = x.node.Children().GetOrElse(list).(g.List)
					children = rec(empty, fromMarksToDepthNode(nodes, depth+1), depth+1)
					z        = children.Children().GetOrElse(list).(g.List)
					node     = g.NewTreeNode(x, z)
				)
				return rec(g.NewCons(node, l), y, depth)
			},
			func() g.Any {
				return g.NewTreeNode(newDepthNode(depth, nothing()), l)
			},
		).(g.Tree)
	}
	return Document{
		List: rec(empty, fromMarksToDepthNode(fromMarks(m), 0), 0),
	}
}
