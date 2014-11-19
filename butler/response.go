package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

type response struct {
	list g.List
}

func Response(list g.List) response {
	return response{
		list: list,
	}
}

func (r response) Build() g.StateT {
	var (
		x = g.StateT_.Of(g.Writer_.Of(g.Empty{}))
		y = r.list.FoldLeft(x, func(x g.Any, y g.Any) g.Any {
			return g.AsStateT(x).Chain(g.Get()).
				Chain(g.Merge(AsBuild(y).Build()))
		})
	)
	return g.AsStateT(y)
}
