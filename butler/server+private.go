package butler

import (
	"net/http"

	g "github.com/SimonRichardson/butler/generic"
)

var (
	noContent       = g.Empty{}
	notFoundService = func(r *http.Request) func() g.Any {
		return func() g.Any {
			// We build the not found service at run time.
			var (
				request  = Request()
				response = Response().ContentType(r.Header.Get("content-type"))
			)
			return Service(request, response, func() g.Any {
				return error404()
			})
		}
	}
	redirectService = func(r *http.Request) func(x g.Any) g.Any {
		return func(x g.Any) g.Any {
			// We build the redirect service at run time.
			var (
				request  = Request()
				response = Response().ContentType(r.Header.Get("content-type"))
			)
			return Service(request, response, func() g.Any {
				return noContent
			})
		}
	}
)

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
