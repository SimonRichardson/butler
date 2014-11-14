package markdown

import (
	"fmt"

	g "github.com/SimonRichardson/butler/generic"
)

type depthNode struct {
	depth int
	node  mark
}

func newDepthNode(depth int, node mark) depthNode {
	return depthNode{
		depth: depth,
		node:  node,
	}
}

func toMarks(s g.List) []mark {
	return s.FoldLeft([]mark{}, func(a, b g.Any) g.Any {
		return append(a.([]mark), b.(mark))
	}).([]mark)
}

func fromMarks(s []mark) g.List {
	var rec func(g.List, []mark) g.List
	rec = func(l g.List, v []mark) g.List {
		num := len(v)
		if num < 1 {
			return l
		}
		return rec(g.NewCons(v[num-1], l), v[:num-1])
	}
	return rec(g.Nil{}, s)
}

func fromMarksToDepthNode(l g.List, depth int) g.List {
	return l.Map(func(x g.Any) g.Any {
		return newDepthNode(depth, x.(mark))
	})
}

func indent(amount int) string {
	list := g.List_.FromAmount(amount)
	return list.Map(func(x g.Any) g.Any {
		return "    "
	}).FoldLeft("", func(a, b g.Any) g.Any {
		return fmt.Sprintf("%s%s", a, b)
	}).(string)
}

func emptyList() g.List {
	return g.List_.Empty()
}

func children(o g.Option) g.List {
	list := func() g.Any {
		return emptyList()
	}
	return g.AsList(o.GetOrElse(list))
}
