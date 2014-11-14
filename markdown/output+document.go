package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type Document struct {
	List g.Tree
}

func (d Document) IsBlock() bool {
	return true
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
		if y.node.IsBlock() {
			return fmt.Sprintf("%s\n%s\n%s", x, z, indent(y.depth))
		} else {
			return fmt.Sprintf("%s%s", x, z)
		}
	}).(string)
}

func document(m ...marks) Document {
	var rec func(g.List, g.List, int) g.Tree
	rec = func(l g.List, m g.List, depth int) g.Tree {
		return m.Head().Fold(
			func(a g.Any) g.Any {
				var (
					x     = a.(depthNode)
					y     = m.Tail()
					nodes = children(x.node.Children())
					tree  = rec(emptyList(), fromMarksToDepthNode(nodes, depth+1), depth+1)
					leaf  = g.NewTreeNode(x, children(tree.Children()))
				)
				return rec(g.NewCons(leaf, l), y, depth)
			},
			func() g.Any {
				return g.NewTreeNode(newDepthNode(depth, nothing()), l)
			},
		).(g.Tree)
	}
	return Document{
		List: rec(emptyList(), fromMarksToDepthNode(fromMarks(m), 0), 0),
	}
}
