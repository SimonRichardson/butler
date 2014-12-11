package butler

import g "github.com/SimonRichardson/butler/generic"

func concat(a, b Server) Server {
	return Server{
		routes: g.Walker_.Combine(a.routes, b.routes, func(a, b g.Any) g.Option {
			var (
				x = g.AsTuple2(a)
				y = g.AsTuple2(b)
			)
			if x.Fst() == y.Fst() {
				sum := g.AsList(x.Snd()).Concat(g.AsList(y.Snd()))
				return g.Option_.Of(g.NewTuple2(x.Fst(), sum))
			}
			return g.Option_.Empty()
		}),
		routeList: a.routeList.Concat(b.routeList),
	}
}
