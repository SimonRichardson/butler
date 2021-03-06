package butler

import (
	g "github.com/SimonRichardson/butler/generic"
	h "github.com/SimonRichardson/butler/http"
)

func flatten(a g.Tuple2) g.List {
	var rec func(l g.List, t g.Tuple2) g.List
	rec = func(l g.List, t g.Tuple2) g.List {
		if b, ok := t.Fst().(g.Tuple2); ok {
			return rec(
				g.NewCons(t.Snd(), l),
				b,
			)
		} else {
			return g.NewCons(t.Snd(), l)
		}
	}
	return rec(g.List_.Empty(), a)
}

func getRoute(a g.List) g.Option {
	return a.Find(func(a g.Any) bool {
		_, ok := a.(h.Route)
		return ok
	})
}
