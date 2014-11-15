package butler

import g "github.com/SimonRichardson/butler/generic"

type request struct {
	list g.List
}

func Request(list g.List) request {
	return request{
		list: list,
	}
}

func (r request) Build() g.StateT {
	var (
		x = g.StateT_.Of(g.Writer_.Of(g.Empty{}))
		y = r.list.FoldLeft(x, func(x g.Any, y g.Any) g.Any {
			return g.AsStateT(x).Chain(g.Get()).
				Chain(g.Merge(AsBuild(y).Build()))
		})
	)
	return g.AsStateT(y)
}
