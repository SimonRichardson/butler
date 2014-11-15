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
	return g.AsStateT(r.list.FoldLeft(g.StateT_.Of(""), func(x g.Any, y g.Any) g.Any {
		return g.AsStateT(x).Chain(g.Get()).
			Chain(g.Merge(AsBuild(y).Build()))
	}))
}
