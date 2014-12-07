package butler

import g "github.com/SimonRichardson/butler/generic"

func concat(a, b Server) Server {
	return Server{
		routes: g.Tree_.Walker(a.routes).Merge(b.routes, func(a, b g.Any) bool {
			var (
				x = g.AsTuple2(a)
				y = g.AsTuple2(b)
			)
			return x.Fst() == y.Fst()
		}),
		routeList: a.routeList.Concat(b.routeList),
	}
}
