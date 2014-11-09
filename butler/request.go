package butler

import (
	g "github.com/SimonRichardson/butler/generic"
)

type request struct {
	list g.List
}

func Request(list g.List) request {
	return request{
		list: list,
	}
}

func (r request) Build() g.StateT {
	return g.AsStateT(r.list.FoldLeft(g.StateT_.Of(""), func(x g.Any, y g.Any) g.Any {
		return g.AsStateT(x).Chain(g.Get()).
			Chain(g.Merge(asBuild(y).Build()))
	}))
}
